package pulse

import (
	"errors"
	"net"

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

func (c *Conn) writeRoute(id int32, group string, routeMode packet.RouteMode, msg *packet.Msg) error {
	if msg == nil {
		return errors.New("msg is nil")
	}
	p := new(packet.Packet)
	p.Udid = c.udid
	p.Type = packet.Type_Body
	p.Msg = msg

	b, err := Encode(p)
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

func (c *Conn) Write(msg *packet.Msg) error {
	return c.writeRoute(0, "", packet.RouteMode_Not, msg)
}

func (c *Conn) WriteRoute(id int32, msg *packet.Msg) error {
	return c.writeRoute(id, "pulse", packet.RouteMode_Normal, msg)
}

func (c *Conn) WriteRouteGroup(id int32, group string, msg *packet.Msg) error {
	return c.writeRoute(id, group, packet.RouteMode_Normal, msg)
}

func (c *Conn) WriteDynamic(id int32, msg *packet.Msg) error {
	return c.writeRoute(id, "", packet.RouteMode_Dynamic, msg)
}

func (c *Conn) WriteDynamicGroup(id int32, group string, msg *packet.Msg) error {
	return c.writeRoute(id, group, packet.RouteMode_Dynamic, msg)
}
