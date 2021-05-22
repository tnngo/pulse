package client

import (
	"bufio"
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"net"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/tnngo/log"
	"github.com/tnngo/pulse"
	"github.com/tnngo/pulse/ip"
	"github.com/tnngo/pulse/packet"
	"github.com/tnngo/pulse/route"
)

type (
	callConnectFunc func() *packet.Msg
	callConnAckFunc func(context.Context, *packet.Msg)
	callCloseFunc   func(context.Context)
)

type Client struct {
	network, addr string

	heartRate, readTimeOut time.Duration

	netconn net.Conn

	udid     string
	secret   string
	authMode packet.AuthMode

	callConnectFunc callConnectFunc
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

func (c *Client) Dial() {
	if c.heartRate == 0 {
		c.heartRate = 30 * time.Second
	}

	if c.readTimeOut == 0 {
		c.readTimeOut = c.heartRate + 30*time.Second
	}
	go c.dial()
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
	if c.callConnectFunc != nil {
		p.Msg = c.callConnectFunc()
	}
	if c.authMode == packet.AuthMode_HmacSha1 {
		p.Secret = c.secret
	}

	p.LocalAddr, _ = ip.GetLocalIP()

	p.Udid = c.udid
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
		netconn.SetReadDeadline(time.Now().Add(c.readTimeOut))
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
					log.L().Info("successfully connected to the server")
					if c.callConnAckFunc != nil {
						c.callConnAckFunc(context.Background(), p.Msg)
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
	if p.Msg == nil {
		log.L().Warn("packet.Msg is nil")
		return
	}

	if p.Route == nil {
		log.L().Warn("packet.Route is nil")
		return
	}

	if p.Route.Group == "pulse" {
		rf, err := route.GetRoute(p.Route.Id)
		if err != nil {
			log.L().Error(err.Error())
			return
		}
		rf(context.Background(), p.Msg)
		return
	}

	f, err := route.GetRouteGroup(p.Route.Id, p.Route.Group)
	if err != nil {
		log.L().Error(err.Error())
		return
	}

	f(context.Background(), p.Msg)
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

func (c *Client) writeMsg(msg *packet.Msg) error {
	if c.getConn() == nil {
		return errors.New("No connection available, the connection object is nil")
	}

	p := new(packet.Packet)
	p.Type = packet.Type_Body
	p.Msg = msg

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

// UDID set the unique identifier of the client.
func (c *Client) UDID(udid string) {
	c.udid = udid
}

func (c *Client) Secret(key, value string) {
	hsha1 := hmac.New(sha1.New, []byte(key+"."+value))
	hsha1.Write([]byte(key + "." + value + "." + c.udid))
	result := hsha1.Sum(nil)
	c.secret = base64.StdEncoding.EncodeToString(result)
}

func (c *Client) CallConnect(f callConnectFunc) {
	c.callConnectFunc = f
}

func (c *Client) CallConnAck(f callConnAckFunc) {
	c.callConnAckFunc = f
}

func (c *Client) WriteMsg(msg *packet.Msg) error {
	return c.writeMsg(msg)
}
