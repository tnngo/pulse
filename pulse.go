package pulse

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"github.com/tnngo/log"
	"github.com/tnngo/pulse/cache"
	"github.com/tnngo/pulse/conn"
	"github.com/tnngo/pulse/ip"
	"github.com/tnngo/pulse/packet"
	"github.com/tnngo/pulse/route"
	"go.uber.org/zap"
)

const (
	CTX_CONN   = "conn"
	CTX_REQ_ID = "request_id"
)

type (
	callConnectFunc func(context.Context, []byte) []byte
	callCloseFunc   func(context.Context)
)

var (
	_connCache = cache.New()
)

type pulse struct {
	network     string
	port        int
	readTimeOut time.Duration

	// buffer pool.
	bufferPool *sync.Pool
	// packet object pool.
	packetPool *sync.Pool

	callConnectFunc callConnectFunc
	callCloseFunc   callCloseFunc
}

// network: tcp、udp
func New(network string, port int) *pulse {
	return newPulse(network, port)
}

func newPulse(network string, port int) *pulse {
	if log.L() == nil {
		log.NewSimple()
	}
	pulse := &pulse{
		network: network,
		port:    port,
		bufferPool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
		packetPool: &sync.Pool{
			New: func() interface{} {
				return new(packet.Packet)
			},
		},
	}
	return pulse
}

func (pl *pulse) ReadTimeOut(t time.Duration) {
	pl.readTimeOut = t
}

func (pl *pulse) Listen() error {
	if pl.readTimeOut == 0 {
		pl.readTimeOut = 60 * time.Second
	}
	switch pl.network {
	case "tcp", "tcp4", "tcp6":
		return pl.listen()
	case "udp", "udp4", "udp6":
		// TODO
		return nil
	default:
		return errors.New("Unsupported network protocol: " + pl.network)
	}
}

func (pl *pulse) listen() error {
	addr := fmt.Sprintf(":%d", pl.port)
	ln, err := net.Listen(pl.network, addr)
	if err != nil {
		return err
	}

	ipAddr, err := ip.GetLocalIP()
	if err != nil {
		return err
	}
	localAddr := fmt.Sprintf("%s:%d", ipAddr, pl.port)

	log.L().Info("Service started successfully", zap.String("listen", localAddr), zap.String("network", pl.network))

	pl.accept(ln)
	return nil
}

// receive tcp connection and message.
func (s *pulse) accept(ln net.Listener) {
	for {
		tcpConn, err := ln.Accept()
		if err != nil {
			log.L().Error(err.Error())
		}

		go s.handle(tcpConn)
	}
}

// handling connections and messages。
func (pl *pulse) handle(netconn net.Conn) {
	reader := bufio.NewReader(netconn)
	buf := pl.bufferPool.Get().(*bytes.Buffer)
	var (
		offset    int
		varintLen int
		size      int

		ctx context.Context
	)
	for {
		// set timeout.
		netconn.SetReadDeadline(time.Now().Add(pl.readTimeOut))
		for {
			b, err := reader.ReadByte()
			if err != nil {
				log.L().Warn(err.Error())
				netconn.Close()
				if ctx == nil {
					return
				}

				ctxConn := GetCtxConn(ctx)
				if mapConn := _connCache.Get(ctxConn.Udid); mapConn == nil {
					return
				} else {
					if ctxConn.ConnectTime < mapConn.ConnectTime {
						return
					}
					_connCache.Del(ctxConn.Udid)
				}

				if pl.callCloseFunc != nil {
					pl.callCloseFunc(ctx)
				}
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
				err := proto.Unmarshal(buf.Next(offset), p)
				if err != nil {
					log.L().Error(err.Error())
				}

				if p.Type == packet.Type_Connect {
					// type connect,
					// connect type is handled separately.
					ctx = pl.connect(netconn, p)
					if pl.callConnectFunc != nil {
						pAck := new(packet.Packet)
						pAck.Type = packet.Type_ConnAck
						pAck.Udid = p.Udid

						repBody := pl.callConnectFunc(ctx, p.Body)
						if repBody != nil {
							// type connack,
							// connack type is handled separately.
							pAck.Body = repBody
						}
						rbyte, err := Encode(pAck)
						if err != nil {
							log.L().Error(err.Error())
							netconn.Close()
							return
						}
						netconn.Write(rbyte)
					}
				} else {
					pl.parse(ctx, netconn, p)
				}

				buf.Reset()
				offset, varintLen, size = 0, 0, 0
				break
			}
		}
	}
}

func (pl *pulse) connect(netconn net.Conn, p *packet.Packet) context.Context {
	c := conn.New(netconn)
	if len(p.Udid) == 0 {
		c.Udid = uuid.New().String()
		p.Udid = c.Udid
	} else {
		if conn := _connCache.Get(p.Udid); conn != nil {
			_connCache.Del(p.Udid)
			if pl.callCloseFunc != nil {
				oldctx := pl.setCtxConn(context.Background(), conn)
				pl.callCloseFunc(oldctx)
			}
			conn.GetNetConn().Close()
			time.Sleep(5 * time.Second)
		}
		c.Udid = p.Udid
	}
	c.Network = netconn.RemoteAddr().Network()
	c.LocalAddr = p.LocalAddr
	c.RemoteAddr = netconn.RemoteAddr().String()
	c.ConnectTime = time.Now().UnixNano() / 1e6
	_connCache.Put(c.Udid, c)
	ctx := pl.setCtxConn(context.Background(), c)
	return ctx
}

// set connection's context.
func (pl *pulse) setCtxConn(ctx context.Context, c *conn.Conn) context.Context {
	return context.WithValue(ctx, CTX_CONN, c)
}

func (pl *pulse) setCtxReqId(ctx context.Context, reqId string) context.Context {
	return context.WithValue(ctx, CTX_REQ_ID, reqId)
}

func (pl *pulse) parse(ctx context.Context, netconn net.Conn, p *packet.Packet) {
	switch p.Type {
	case packet.Type_Ping:
		pl.pong(netconn)
	case packet.Type_Body:
		pl.body(ctx, p)
	}
}

func (pl *pulse) pong(netconn net.Conn) {
	wp := new(packet.Packet)
	wp.Type = packet.Type_Pong
	b, err := Encode(wp)
	if err != nil {
		log.L().Error(err.Error())
		return
	}
	netconn.Write(b)
}

func (pl *pulse) body(ctx context.Context, p *packet.Packet) {
	f, err := route.Get(p.RouteId)
	if err != nil {
		log.L().Error(err.Error())
		return
	}
	f(ctx, p.Body)
}

func (pl *pulse) CallConnect(f callConnectFunc) {
	pl.callConnectFunc = f
}

func (pl *pulse) CallClose(f callCloseFunc) {
	pl.callCloseFunc = f
}

// Encode encapsulate protobuf.
func Encode(p *packet.Packet) ([]byte, error) {
	b1, err := proto.Marshal(p)
	if err != nil {
		return nil, err
	}

	b2 := proto.EncodeVarint(uint64(len(b1)))
	b3 := []byte{byte(len(b2))}
	b3 = append(b3, b2...)
	b3 = append(b3, b1...)
	return b3, nil
}

func GetCtxConn(ctx context.Context) *conn.Conn {
	return ctx.Value(CTX_CONN).(*conn.Conn)
}

func GetCtxRequestId(ctx context.Context) string {
	return ctx.Value(CTX_REQ_ID).(string)
}

func GetConn(udid string) *conn.Conn {
	return _connCache.Get(udid)
}
