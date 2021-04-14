package conn

import (
	"net"

	"github.com/golang/protobuf/proto"
	"github.com/tnngo/pulse/packet"
)

type Conn struct {
	// network connection.
	netconn net.Conn

	// client unique identifier.
	Udid    string
	Network string
	// client remote address.
	RemoteAddr string
	// client local address.
	LocalAddr string
	// client connection time.
	ConnectTime int64
}

func NewConn(netconn net.Conn) *Conn {
	return &Conn{
		netconn: netconn,
	}
}

func (c *Conn) writeRoute(id int32, body []byte) error {
	p := new(packet.Packet)
	p.Udid = c.Udid
	p.Type = packet.Type_Body
	p.RouteId = id
	p.LocalAddr = c.LocalAddr
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

func (c *Conn) WriteRoute(id int32, body []byte) error {
	return c.writeRoute(id, body)
}

func (c *Conn) GetNetConn() net.Conn {
	return c.netconn
}
