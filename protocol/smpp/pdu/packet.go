package pdu

import (
	"github.com/sujit-baniya/smpp/coding"
)

type Responsable interface {
	Resp() interface{}
}

type Headerable interface {
	GetHeader() Header
}

// AlertNotification see SMPP v5, section 4.1.3.1 (64p)
type AlertNotification struct {
	Header     Header `id:"00000102"`
	SourceAddr Address
	ESMEAddr   Address
	Tags       Tags
}

func (p *AlertNotification) GetHeader() Header {
	return p.Header
}

// BindReceiver see SMPP v5, section 4.1.1.3 (58p)
type BindReceiver struct {
	Header       Header `id:"00000001"`
	SystemID     string
	Password     string
	SystemType   string
	Version      InterfaceVersion
	AddressRange Address // see section 4.7.3.1
}

func (p *BindReceiver) GetHeader() Header {
	return p.Header
}

func (p *BindReceiver) Resp() interface{} {
	return &BindReceiverResp{Header: Header{Sequence: p.Header.Sequence}, SystemID: p.SystemID}
}

// BindReceiverResp see SMPP v5, section 4.1.1.4 (59p)
type BindReceiverResp struct {
	Header   Header `id:"80000001"`
	SystemID string
	Tags     Tags
}

func (p *BindReceiverResp) GetHeader() Header {
	return p.Header
}

// BindTransceiver see SMPP v5, section 4.1.1.5 (59p)
type BindTransceiver struct {
	Header       Header `id:"00000009"`
	SystemID     string
	Password     string
	SystemType   string
	Version      InterfaceVersion
	AddressRange Address // see section 4.7.3.1
}

func (p *BindTransceiver) GetHeader() Header {
	return p.Header
}

func (p *BindTransceiver) Resp() interface{} {
	return &BindTransceiverResp{Header: Header{Sequence: p.Header.Sequence}, SystemID: p.SystemID}
}

// BindTransceiverResp see SMPP v5, section 4.1.1.6 (60p)
type BindTransceiverResp struct {
	Header   Header `id:"80000009"`
	SystemID string
	Tags     Tags
}

func (p *BindTransceiverResp) GetHeader() Header {
	return p.Header
}

// BindTransmitter see SMPP v5, section 4.1.1.1 (56p)
type BindTransmitter struct {
	Header       Header `id:"00000002"`
	SystemID     string
	Password     string
	SystemType   string
	Version      InterfaceVersion
	AddressRange Address // see section 4.7.3.1
}

func (p *BindTransmitter) GetHeader() Header {
	return p.Header
}

func (p *BindTransmitter) Resp() interface{} {
	return &BindTransmitterResp{Header: Header{Sequence: p.Header.Sequence}, SystemID: p.SystemID}
}

// BindTransmitterResp see SMPP v5, section 4.1.1.2 (57p)
type BindTransmitterResp struct {
	Header   Header `id:"80000002"`
	SystemID string
	Tags     Tags
}

func (p *BindTransmitterResp) GetHeader() Header {
	return p.Header
}

// BroadcastSM see SMPP v5, section 4.4.1.1 (92p)
type BroadcastSM struct {
	Header               Header `id:"00000112"`
	ServiceType          string
	MessageID            string
	ScheduleDeliveryTime string
	ValidityPeriod       string
	SourceAddr           Address
	DataCoding           coding.DataCoding
	ReplaceIfPresent     bool
	PriorityFlag         byte
	DefaultMessageID     byte
	Tags                 Tags
}

func (p *BroadcastSM) GetHeader() Header {
	return p.Header
}

func (p *BroadcastSM) Resp() interface{} {
	return &BroadcastSMResp{Header: Header{Sequence: p.Header.Sequence}, MessageID: p.MessageID}
}

// BroadcastSMResp see SMPP v5, section 4.4.1.2 (96p)
type BroadcastSMResp struct {
	Header    Header `id:"80000112"`
	MessageID string
	Tags      Tags
}

func (p *BroadcastSMResp) GetHeader() Header {
	return p.Header
}

// CancelBroadcastSM see SMPP v5, section 4.6.2.1 (110p)
type CancelBroadcastSM struct {
	Header      Header `id:"00000113"`
	ServiceType string
	MessageID   string
	SourceAddr  Address
	Tags        Tags
}

func (p *CancelBroadcastSM) GetHeader() Header {
	return p.Header
}

func (p *CancelBroadcastSM) Resp() interface{} {
	return &CancelBroadcastSMResp{Header: Header{Sequence: p.Header.Sequence}}
}

// CancelBroadcastSMResp see SMPP v5, section 4.6.2.3 (112p)
type CancelBroadcastSMResp struct {
	Header Header `id:"80000113"`
}

func (p *CancelBroadcastSMResp) GetHeader() Header {
	return p.Header
}

// CancelSM see SMPP v5, section 4.5.1.1 (100p)
type CancelSM struct {
	Header      Header `id:"00000008"`
	ServiceType string
	MessageID   string
	SourceAddr  Address
	DestAddr    Address
}

func (p *CancelSM) GetHeader() Header {
	return p.Header
}

func (p *CancelSM) Resp() interface{} {
	return &CancelSMResp{Header: Header{Sequence: p.Header.Sequence}}
}

// CancelSMResp see SMPP v5, section 4.5.1.2 (101p)
type CancelSMResp struct {
	Header Header `id:"80000008"`
}

func (p *CancelSMResp) GetHeader() Header {
	return p.Header
}

// DataSM see SMPP v5, section 4.2.2.1 (69p)
type DataSM struct {
	Header             Header `id:"00000103"`
	ServiceType        string
	SourceAddr         Address
	DestAddr           Address
	ESMClass           ESMClass
	RegisteredDelivery RegisteredDelivery
	DataCoding         coding.DataCoding
	Tags               Tags
}

func (p *DataSM) GetHeader() Header {
	return p.Header
}

func (p *DataSM) Resp() interface{} {
	return &DataSMResp{Header: Header{Sequence: p.Header.Sequence}}
}

// DataSMResp see SMPP v5, section 4.2.2.2 (70p)
type DataSMResp struct {
	Header    Header `id:"80000103"`
	MessageID string
	Tags      Tags
}

func (p *DataSMResp) GetHeader() Header {
	return p.Header
}

// DeliverSM see SMPP v5, section 4.3.1.1 (85p)
type DeliverSM struct {
	Header               Header `id:"00000005"`
	ServiceType          string
	SourceAddr           Address
	DestAddr             Address
	ESMClass             ESMClass
	ProtocolID           byte
	PriorityFlag         byte
	ScheduleDeliveryTime string
	ValidityPeriod       string
	RegisteredDelivery   RegisteredDelivery
	ReplaceIfPresent     bool
	Message              ShortMessage
	Tags                 Tags
}

func (p *DeliverSM) GetHeader() Header {
	return p.Header
}

func (p *DeliverSM) Resp() interface{} {
	return &DeliverSMResp{Header: Header{Sequence: p.Header.Sequence}}
}

// DeliverSMResp see SMPP v5, section 4.3.1.1 (87p)
type DeliverSMResp struct {
	Header    Header `id:"80000005"`
	MessageID string
	Tags      Tags
}

func (p *DeliverSMResp) GetHeader() Header {
	return p.Header
}

// EnquireLink see SMPP v5, section 4.1.2.1 (63p)
type EnquireLink struct {
	Header Header `id:"00000015"`
	Tags   Tags
}

func (p *EnquireLink) GetHeader() Header {
	return p.Header
}

func (p *EnquireLink) Resp() interface{} {
	return &EnquireLinkResp{Header: Header{Sequence: p.Header.Sequence}}
}

// EnquireLinkResp see SMPP v5, section 4.1.2.2 (63p)
type EnquireLinkResp struct {
	Header Header `id:"80000015"`
}

func (p *EnquireLinkResp) GetHeader() Header {
	return p.Header
}

// GenericNACK see SMPP v5, section 4.1.4.1 (65p)
type GenericNACK struct {
	Header Header `id:"80000000"`
	Tags   Tags
}

func (p *GenericNACK) GetHeader() Header {
	return p.Header
}

// Outbind see SMPP v5, section 4.1.1.7 (61p)
type Outbind struct {
	Header   Header `id:"0000000B"`
	SystemID string
	Password string
}

func (p *Outbind) GetHeader() Header {
	return p.Header
}

// QueryBroadcastSM see SMPP v5, section 4.6.1.1 (107p)
type QueryBroadcastSM struct {
	Header     Header `id:"00000111"`
	MessageID  string
	SourceAddr Address
	Tags       Tags
}

func (p *QueryBroadcastSM) GetHeader() Header {
	return p.Header
}

func (p *QueryBroadcastSM) Resp() interface{} {
	return &QueryBroadcastSMResp{Header: Header{Sequence: p.Header.Sequence}, MessageID: p.MessageID}
}

// QueryBroadcastSMResp see SMPP v5, section 4.6.1.3 (108p)
type QueryBroadcastSMResp struct {
	Header    Header `id:"80000111"`
	MessageID string
	Tags      Tags
}

func (p *QueryBroadcastSMResp) GetHeader() Header {
	return p.Header
}

// QuerySM see SMPP v5, section 4.5.2.1 (101p)
type QuerySM struct {
	Header     Header `id:"00000003"`
	MessageID  string
	SourceAddr Address
}

func (p *QuerySM) GetHeader() Header {
	return p.Header
}

func (p *QuerySM) Resp() interface{} {
	return &QuerySMResp{Header: Header{Sequence: p.Header.Sequence}}
}

// QuerySMResp see SMPP v5, section 4.5.2.2 (103p)
type QuerySMResp struct {
	Header       Header `id:"80000003"`
	MessageID    string
	FinalDate    string
	MessageState MessageState
	ErrorCode    CommandStatus
}

func (p *QuerySMResp) GetHeader() Header {
	return p.Header
}

// ReplaceSM see SMPP v5, section 4.5.3.1 (104p)
type ReplaceSM struct {
	Header               Header `id:"00000007"`
	MessageID            string
	SourceAddr           Address
	ScheduleDeliveryTime string
	ValidityPeriod       string
	RegisteredDelivery   RegisteredDelivery
	Message              ShortMessage
	Tags                 Tags
}

func (p *ReplaceSM) GetHeader() Header {
	return p.Header
}

func (p *ReplaceSM) Resp() interface{} {
	return &ReplaceSMResp{Header: Header{Sequence: p.Header.Sequence}}
}

// ReplaceSMResp see SMPP v5, section 4.5.3.2 (106p)
type ReplaceSMResp struct {
	Header Header `id:"80000007"`
}

func (p *ReplaceSMResp) GetHeader() Header {
	return p.Header
}

// SubmitMulti see SMPP v5, section 4.2.3.1 (71p)
type SubmitMulti struct {
	Header               Header `id:"00000021"`
	ServiceType          string
	SourceAddr           Address
	DestAddrList         DestinationAddresses
	ESMClass             ESMClass
	ProtocolID           byte
	PriorityFlag         byte
	ScheduleDeliveryTime string
	ValidityPeriod       string
	RegisteredDelivery   RegisteredDelivery
	ReplaceIfPresent     bool
	Message              ShortMessage
	Tags                 Tags
}

func (p *SubmitMulti) GetHeader() Header {
	return p.Header
}

func (p *SubmitMulti) Resp() interface{} {
	return &SubmitMultiResp{Header: Header{Sequence: p.Header.Sequence}}
}

// SubmitMultiResp see SMPP v5, section 4.2.3.2 (74p)
type SubmitMultiResp struct {
	Header           Header `id:"80000021"`
	MessageID        string
	UnsuccessfulSMEs UnsuccessfulRecords
	Tags             Tags
}

func (p *SubmitMultiResp) GetHeader() Header {
	return p.Header
}

// SubmitSM see SMPP v5, section 4.2.1.1 (66p)
type SubmitSM struct {
	Header               Header `id:"00000004"`
	ServiceType          string
	SourceAddr           Address
	DestAddr             Address
	ESMClass             ESMClass
	ProtocolID           byte
	PriorityFlag         byte
	ScheduleDeliveryTime string
	ValidityPeriod       string
	RegisteredDelivery   RegisteredDelivery
	ReplaceIfPresent     bool
	Message              ShortMessage
	Tags                 Tags
}

func (p *SubmitSM) GetHeader() Header {
	return p.Header
}

func (p *SubmitSM) Resp() interface{} {
	return &SubmitSMResp{Header: Header{Sequence: p.Header.Sequence}}
}

// SubmitSMResp see SMPP v5, section 4.2.1.2 (68p)
type SubmitSMResp struct {
	Header    Header `id:"80000004"`
	MessageID string
}

func (p *SubmitSMResp) GetHeader() Header {
	return p.Header
}

// Unbind see SMPP v5, section 4.1.1.8 (61p)
type Unbind struct {
	Header Header `id:"00000006"`
}

func (p *Unbind) GetHeader() Header {
	return p.Header
}

func (p *Unbind) Resp() interface{} {
	return &UnbindResp{Header: Header{Sequence: p.Header.Sequence}}
}

// UnbindResp see SMPP v5, section 4.1.1.9 (62p)
type UnbindResp struct {
	Header Header `id:"80000006"`
}

func (p *UnbindResp) GetHeader() Header {
	return p.Header
}
