package pulse

import (
	"net"

	"github.com/golang/protobuf/proto"
	"github.com/tnngo/pulse/packet"
)

type Conn struct {
	// network connection.
	netconn net.Conn

	// client unique identifier.
	udid    string
	network string
	// client remote address.
	remoteAddr string
	// client local address.
	localAddr string
	// client connection time.
	connectTime int64
}

func newConn(netconn net.Conn) *Conn {
	return &Conn{
		netconn: netconn,
	}
}

func (c *Conn) writeRoute(id int32, body []byte) error {
	p := new(packet.Packet)
	p.Udid = c.udid
	p.Type = packet.Type_Body
	p.RouteId = id
	p.LocalAddr = c.localAddr
	p.Body = body

	b, err := proto.Marshal(p)
	if err != nil {
		return err
	}

	_, err = c.netconn.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func (c *Conn) UDID() string {
	return c.udid
}

func (c *Conn) Network() string {
	return c.network
}

func (c *Conn) RemoteAddr() string {
	return c.remoteAddr
}

func (c *Conn) LocalAddr() string {
	return c.localAddr
}

func (c *Conn) ConnectTime() int64 {
	return c.connectTime
}

func (c *Conn) WriteRoute(id int32, body []byte) error {
	return c.writeRoute(id, body)
}
