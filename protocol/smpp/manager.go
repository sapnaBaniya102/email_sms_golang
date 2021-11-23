package smpp

import (
	"context"
	"errors"
	"github.com/rs/xid"
	"github.com/sujit-baniya/smpp/balancer"
	"github.com/sujit-baniya/smpp/coding"
	"github.com/sujit-baniya/smpp/pdu"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"
)

type ManagerInterface interface {
	Start() error
	AddConnection(noOfConnection ...int) error
	RemoveConnection(connectionID ...string) error
	GetConnection(conIds ...string) ConnectionInterface
	SetupConnection() error
	Rebind() error
	Send(payload interface{}, connectionID ...string) (interface{}, error)
	Close(connectionID ...string) error
}

type Auth struct {
	SystemID   string
	Password   string
	SystemType string
}

type Setting struct {
	Name             string
	Slug             string
	URL              string
	Auth             Auth
	SmppVersion      pdu.InterfaceVersion
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	EnquiryInterval  time.Duration
	EnquiryTimeout   time.Duration
	MaxConnection    int
	Balancer         balancer.Balancer
	Throttle         int
	UseAllConnection bool
	HandlePDU        func(conn *Conn)
	AutoRebind       bool
}

type Manager struct {
	Name        string
	Slug        string
	ID          string
	ctx         context.Context
	setting     Setting
	connections map[string]*Conn
	Balancer    balancer.Balancer
	connIDs     []string
	mu          sync.RWMutex
}

type HandlePDU func(conn *Conn)

type Message struct {
	From    string
	To      string
	Message string
}

func NewManager(setting Setting) (*Manager, error) {
	if setting.MaxConnection == 0 {
		setting.MaxConnection = 1
	}
	manager := &Manager{
		Name:        setting.Name,
		Slug:        setting.Slug,
		ID:          xid.New().String(),
		ctx:         context.Background(),
		setting:     setting,
		connections: make(map[string]*Conn),
	}
	if setting.Balancer == nil {
		manager.Balancer = &balancer.RoundRobin{}
	}
	return manager, nil
}

func (m *Manager) Start() error {
	if m.setting.UseAllConnection {
		for i := 0; i < m.setting.MaxConnection; i++ {
			err := m.SetupConnection()
			if err != nil {
				return err
			}
		}
		return nil
	}
	if len(m.connIDs) == 0 {
		err := m.SetupConnection()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) AddConnection(noOfConnection ...int) error {
	con := 1
	if len(noOfConnection) > 0 {
		con = noOfConnection[0]
	}
	if con > m.setting.MaxConnection {
		return errors.New("Can't create more than allowed no of connections.")
	}
	if (len(m.connIDs) + con) > m.setting.MaxConnection {
		return errors.New("There are active sessions. Can't create more than allowed no of sessions.")
	}
	connLeft := m.setting.MaxConnection - len(m.connIDs)
	n := 0
	if connLeft >= con {
		n = con
	} else {
		n = connLeft
	}
	for i := 0; i < n; i++ {
		err := m.SetupConnection()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) RemoveConnection(conID ...string) error {
	if len(conID) > 0 {
		for _, id := range conID {
			if con, ok := m.connections[id]; ok {
				err := con.Close()
				if err != nil {
					return err
				}
				m.connIDs = remove(m.connIDs, id)
				delete(m.connections, id)
			}
		}
	} else {
		for id, con := range m.connections {
			err := con.Close()
			if err != nil {
				return err
			}
			m.connIDs = remove(m.connIDs, id)
			delete(m.connections, id)
		}
	}
	return nil
}

func (m *Manager) Rebind() error {
	m.Close()
	m.connections = make(map[string]*Conn)
	m.connIDs = []string{}
	m.Start()
	return m.HandlePDU()
}

func (m *Manager) SetupConnection() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	parent, err := net.Dial("tcp", m.setting.URL)
	if err != nil {
		return err
	}
	conn := NewConn(context.Background(), parent, m.setting.Throttle)
	conn.WriteTimeout = m.setting.WriteTimeout
	conn.ReadTimeout = m.setting.ReadTimeout
	go conn.Watch()
	resp, err := conn.Bind(context.Background(), &pdu.BindTransceiver{
		SystemID:   m.setting.Auth.SystemID,
		Password:   m.setting.Auth.Password,
		SystemType: m.setting.Auth.SystemType,
		Version:    m.setting.SmppVersion,
	})
	if err != nil {
		return err
	}
	r := resp.(*pdu.BindTransceiverResp)
	if r.Header.CommandStatus == 0 {
		// start keep-alive
		go conn.EnquireLink(m.setting.EnquiryInterval, m.setting.EnquiryTimeout)
		m.connIDs = append(m.connIDs, conn.ID)
		m.connections[conn.ID] = conn
	}
	return nil
}

func (m *Manager) GetConnection(conIds ...string) ConnectionInterface {
	var pickedID string
	if len(conIds) > 0 { // pick among custom
		pickedID, _ = m.Balancer.Pick(conIds)
		if con, ok := m.connections[pickedID]; ok {
			return con
		}
	}

	// pick among managing session
	pickedID, _ = m.Balancer.Pick(m.connIDs)
	con, _ := m.connections[pickedID]
	return con
}

func (m *Manager) Send(payload interface{}, connectionId ...string) (interface{}, error) {
	sms := payload.(Message)
	shortMessages, err := m.Compose(sms.Message)
	if err != nil {
		panic(err)
	}
	responses := make(map[*pdu.SubmitSM]*pdu.SubmitSMResp)
	responseChan := make(chan map[*pdu.SubmitSM]*pdu.SubmitSMResp)
	wg := &sync.WaitGroup{}
	for _, shortMessage := range shortMessages {
		wg.Add(1)
		go m.SendShortMessage(sms.From, sms.To, shortMessage, wg, responseChan, connectionId...)
	}
	go func() {
		wg.Wait()
		close(responseChan)
	}()
	for response := range responseChan {
		for submitSM, submitSMResp := range response {
			responses[submitSM] = submitSMResp
		}
	}
	return responses, nil
}

func (m *Manager) SendShortMessage(from string, to string, shortMessage pdu.ShortMessage, wg *sync.WaitGroup, responseChan chan<- map[*pdu.SubmitSM]*pdu.SubmitSMResp, connectionId ...string) error {

	defer wg.Done()
	conn := m.GetConnection(connectionId...).(*Conn)
	packet := m.Prepare(from, to, shortMessage)
	err := conn.Throttle()
	if err != nil {
		return err
	}
	resp, err := conn.Submit(m.ctx, packet)
	if err != nil {
		return err
	}
	mp := map[*pdu.SubmitSM]*pdu.SubmitSMResp{
		packet: resp.(*pdu.SubmitSMResp),
	}
	responseChan <- mp
	return nil
}

func (m *Manager) Prepare(from string, to string, shortMessage pdu.ShortMessage) *pdu.SubmitSM {
	return &pdu.SubmitSM{
		SourceAddr: parseSrcPhone(from),
		DestAddr:   parseDestPhone(to),
		ESMClass:   pdu.ESMClass{UDHIndicator: true},
		RegisteredDelivery: pdu.RegisteredDelivery{
			MCDeliveryReceipt:           1,
			SMEOriginatedAcknowledgment: 1,
			IntermediateNotification:    true,
			Reserved:                    7,
		},
		Message: shortMessage,
	}
}

func (m *Manager) Close(connectionId ...string) error {
	if len(connectionId) > 0 {
		if con, ok := m.connections[connectionId[0]]; ok {
			err := con.Close()
			if err != nil {
				return err
			}
		}
	} else {
		for _, conn := range m.connections {
			err := conn.Close()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *Manager) HandlePDU() error {
	for _, conn := range m.connections {
		go m.setting.HandlePDU(conn)
	}
	return nil
}

func (m *Manager) Compose(msg string) ([]pdu.ShortMessage, error) {
	return Compose(msg, m.setting.SmppVersion)
}

func Compose(msg string, smppVersion pdu.InterfaceVersion) ([]pdu.ShortMessage, error) {
	reference := uint16(rand.Intn(0xFFFF))
	var dataCoding coding.DataCoding
	if smppVersion == pdu.SMPPVersion50 {
		dataCoding = coding.BestCoding(msg)
	} else {
		dataCoding = coding.BestSafeCoding(msg)
	}
	return pdu.ComposeMultipartShortMessage(msg, dataCoding, reference)
}

func parseSrcPhone(phone string) pdu.Address {
	if strings.HasPrefix(phone, "+") {
		return pdu.Address{TON: 1, NPI: 1, No: phone}
	}

	if utf8.RuneCountInString(phone) <= 5 {
		return pdu.Address{TON: 3, NPI: 0, No: phone}
	}
	if isLetter(phone) {
		return pdu.Address{TON: 5, NPI: 0, No: phone}
	}
	return pdu.Address{TON: 1, NPI: 1, No: phone}
}

func parseDestPhone(phone string) pdu.Address {
	if strings.HasPrefix(phone, "+") {
		return pdu.Address{TON: 1, NPI: 1, No: phone}
	}
	return pdu.Address{TON: 0, NPI: 1, No: phone}
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
