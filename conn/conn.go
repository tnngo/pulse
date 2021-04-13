package conn

import (
	"net"
)

type Conn struct {
	// network connection.
	netconn net.Conn

	// client unique identifier.
	UDID    string
	Network string
	// client remote address.
	RemoteAddr string
	// client local address.
	LocalAddr string
	// client connection time.
	ConnectTime int64
}

func New(netconn net.Conn) *Conn {
	return &Conn{
		netconn: netconn,
	}
}

func (c *Conn) GetNetConn() net.Conn {
	return c.netconn
}
