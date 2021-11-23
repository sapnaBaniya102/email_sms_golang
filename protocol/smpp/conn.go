package smpp

import (
	"context"
	"errors"
	"github.com/sujit-baniya/smpp/pdu"
	"golang.org/x/time/rate"
	"io"
	"math/rand"
	"net"
	"time"

	"github.com/rs/xid"
)

type ConnectionInterface interface {
	Send(packet interface{}) (err error)
	Throttle() error
}

type Conn struct {
	parent       net.Conn
	ctx          context.Context
	cancel       context.CancelFunc
	receiveQueue chan interface{}
	pending      map[int32]func(interface{})
	ID           string
	NextSequence func() int32
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	RateLimiter  *rate.Limiter
	rwctx        context.Context
	lmctx        context.Context
}

func OpenConn(ctx context.Context, smsc string, throttle int) (conn *Conn, err error) {
	parent, err := net.Dial("tcp", smsc)
	if err == nil {
		conn = NewConn(ctx, parent, throttle)
	}
	return
}

func NewConn(ctx context.Context, parent net.Conn, throttle int) *Conn {
	ctx, cancel := context.WithCancel(ctx)

	conn := &Conn{
		parent:       parent,
		ctx:          ctx,
		cancel:       cancel,
		receiveQueue: make(chan interface{}),
		pending:      make(map[int32]func(interface{})),
		ID:           xid.New().String(),
		NextSequence: rand.Int31,
		ReadTimeout:  time.Minute * 15,
		WriteTimeout: time.Minute * 15,
	}
	if throttle != 0 {
		rateLimiter := rate.NewLimiter(rate.Limit(throttle), 1)
		conn.rwctx = context.Background()
		conn.RateLimiter = rateLimiter
	}
	return conn
}

//goland:noinspection SpellCheckingInspection
func (c *Conn) Watch() {
	defer c.cancel()
	var err error
	var packet interface{}
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}
		if c.ReadTimeout > 0 {
			_ = c.parent.SetReadDeadline(time.Now().Add(c.ReadTimeout))
		}
		var status pdu.CommandStatus
		if packet, err = pdu.ReadPDU(c.parent); errors.Is(err, io.EOF) {
			return
		} else if ok := errors.Is(err, status); err != nil {
			if packet == nil {
				return
			} else if !ok {
				status = pdu.ErrUnknownError
			}
			sequence := pdu.ReadSequence(packet)
			_ = c.Send(&pdu.GenericNACK{
				Header: pdu.Header{CommandStatus: status, Sequence: sequence},
				Tags:   pdu.Tags{0xFFFF: []byte(err.Error())},
			})
			continue
		} else if callback, ok := c.pending[pdu.ReadSequence(packet)]; ok {
			callback(packet)
		} else {
			c.receiveQueue <- packet
		}
	}
}

func (c *Conn) Bind(ctx context.Context, packet pdu.Responsable) (resp interface{}, err error) {
	return c.Submit(ctx, packet)
}

func (c *Conn) Submit(ctx context.Context, packet pdu.Responsable) (resp interface{}, err error) {
	sequence := c.NextSequence()
	pdu.WriteSequence(packet, sequence)
	if err = c.Send(packet); err != nil {
		return
	}
	returns := make(chan interface{}, 1)
	c.pending[sequence] = func(resp interface{}) { returns <- resp }
	defer delete(c.pending, sequence)
	select {
	case <-c.ctx.Done():
		err = ErrConnectionClosed
	case <-ctx.Done():
		err = ctx.Err()
	case resp = <-returns:
	}
	return
}

func (c *Conn) Send(packet interface{}) (err error) {
	sequence := pdu.ReadSequence(packet)
	if sequence == 0 || sequence < 0 {
		err = pdu.ErrInvalidSequence
		return
	}
	if c.WriteTimeout > 0 {
		err = c.parent.SetWriteDeadline(time.Now().Add(c.WriteTimeout))
	}
	if err == nil {
		_, err = pdu.Marshal(c.parent, packet)
	}
	if errors.Is(err, io.EOF) {
		err = ErrConnectionClosed
	}
	return
}

func (c *Conn) EnquireLink(tick time.Duration, timeout time.Duration) {
	ticker := time.NewTicker(tick)
	defer ticker.Stop()
	sendEnquireLink := func() {
		ctx, cancel := context.WithTimeout(c.ctx, timeout)
		defer cancel()
		if _, err := c.Submit(ctx, new(pdu.EnquireLink)); err != nil {
			ticker.Stop()
			_ = c.Close()
		}
	}
	for {
		sendEnquireLink()
		<-ticker.C
	}
}

func (c *Conn) Close() (err error) {
	ctx, cancel := context.WithTimeout(c.ctx, time.Second)
	defer cancel()
	defer c.cancel()
	if _, err = c.Submit(ctx, new(pdu.Unbind)); err == nil {
		close(c.receiveQueue)
		err = c.parent.Close()
	}
	return
}

func (c *Conn) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *Conn) Throttle() error {
	if c.RateLimiter != nil {
		return c.RateLimiter.Wait(c.rwctx)
	}
	return nil
}

func (c *Conn) PDU() <-chan interface{} {
	return c.receiveQueue
}
