package client

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"github.com/tnngo/log"
	"github.com/tnngo/pulse"
	"github.com/tnngo/pulse/ip"
	"github.com/tnngo/pulse/packet"
	"github.com/tnngo/pulse/route"
)

type (
	callConnAckFunc func(context.Context, *packet.Packet)
	callCloseFunc   func(context.Context)
)

type Client struct {
	network, addr string

	heartRate, readTimeOut time.Duration

	netconn net.Conn

	udid string

	enableReqId bool

	callConnAckFunc callConnAckFunc

	rwmutex sync.RWMutex
}

// New
func New(network string, addr string) *Client {
	if log.L() == nil {
		log.NewSimple()
	}
	return &Client{
		network: network,
		addr:    addr,
	}
}

// UDID set the unique identifier of the client.
func (c *Client) UDID(udid string) {
	c.udid = udid
}

func (c *Client) Dial() {
	if c.heartRate == 0 {
		c.heartRate = 30 * time.Second
	}

	if c.readTimeOut == 0 {
		c.readTimeOut = c.heartRate + 30*time.Second
	}
	c.dial()
}

func (c *Client) dial() {
	for {
		netconn, err := net.Dial(c.network, c.addr)
		if err != nil {
			log.L().Warn(err.Error())
			if netconn != nil {
				netconn.Close()
			}
			time.Sleep(5 * time.Second)
			continue
		}

		if b, err := c.connect(); err != nil {
			log.L().Error(err.Error())
			time.Sleep(5 * time.Second)
			continue
		} else {
			netconn.Write(b)
		}

		c.handle(netconn)
		time.Sleep(5 * time.Second)
	}
}

func (c *Client) connect() ([]byte, error) {
	p := new(packet.Packet)
	la, err := ip.GetLocalIP()
	if err != nil {
		log.L().Error(err.Error())
	}
	p.UDID = c.udid
	p.LocalAddr = la
	p.Type = packet.Type_Connect

	return pulse.Encode(p)
}

func (c *Client) handle(netconn net.Conn) {
	reader := bufio.NewReader(netconn)
	buf := new(bytes.Buffer)
	var (
		offset    int
		varintLen int
		size      int

		//ctx context.Context
	)
	for {
		// set timeout.
		//conn.SetReadDeadline(time.Now().Add(p.readTimeOut))
		netconn.SetReadDeadline(time.Now().Add(3 * time.Minute))
		for {
			b, err := reader.ReadByte()
			if err != nil {
				netconn.Close()
				log.L().Error(err.Error())
				return
			}

			if varintLen == 0 {
				varintLen = int(b)
				continue
			}

			buf.WriteByte(b)
			offset++

			if offset == varintLen {
				px, pn := proto.DecodeVarint(buf.Next(offset))
				size = int(px) + pn
			}

			if offset == size && size != 0 {
				p := new(packet.Packet)
				proto.Unmarshal(buf.Next(offset), p)

				// type connack,
				// connack type is handled separately.
				if p.Type == packet.Type_ConnAck {
					if c.callConnAckFunc != nil {
						c.callConnAckFunc(context.Background(), p)
					}
					c.setConn(netconn)
					go c.heartbeat()
				} else {
					c.parse(p)
				}

				buf.Reset()
				offset, varintLen, size = 0, 0, 0
				break
			}
		}
	}
}

func (c *Client) heartbeat() {
	p := new(packet.Packet)
	p.Type = packet.Type_Ping
	b, _ := pulse.Encode(p)

	for {
		time.Sleep(c.heartRate)
		_, err := c.getConn().Write(b)
		if err != nil {
			c.getConn().Close()
			log.L().Error(err.Error())
			return
		}
	}
}

func (c *Client) parse(p *packet.Packet) {
	switch p.Type {
	case packet.Type_Pong:
		break
	case packet.Type_Body:
		c.body(p)
	}
}

func (c *Client) body(p *packet.Packet) {
	f, err := route.Get(p.RouteId)
	if err != nil {
		log.L().Error(err.Error())
		return
	}
	f(context.Background(), p.Body)
}

func (c *Client) getConn() net.Conn {
	c.rwmutex.RLock()
	defer c.rwmutex.RUnlock()
	return c.netconn
}

func (c *Client) setConn(netconn net.Conn) {
	c.rwmutex.Lock()
	defer c.rwmutex.Unlock()
	c.netconn = netconn
}

func (c *Client) writeRoute(id int32, body []byte) error {
	if c.getConn() == nil {
		return errors.New("No connection available, the connection object is nil")
	}

	p := new(packet.Packet)
	p.LocalAddr = c.getConn().LocalAddr().String()
	if c.enableReqId {
		p.RequestId = strings.Replace(uuid.New().String(), "-", "", -1)
	}
	p.RouteId = id
	p.Type = packet.Type_Body
	p.Body = body
	b, err := pulse.Encode(p)
	if err != nil {
		return err
	}
	_, err = c.getConn().Write(b)
	if err != nil {
		return err
	}
	return nil
}

// EnableRequestId uuid, 32 length,
// however, only WriteRoute method will send RequestId.
func (c *Client) EnableRequestId() {
	c.enableReqId = true
}

func (c *Client) CallConnAck(f callConnAckFunc) {
	c.callConnAckFunc = f
}

//
func (c *Client) WriteRoute(id int32, body []byte) error {
	return c.writeRoute(id, body)
}
