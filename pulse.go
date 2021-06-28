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
	"github.com/tnngo/pulse/ip"
	"github.com/tnngo/pulse/packet"
	"github.com/tnngo/pulse/route"
	"go.uber.org/zap"
)

var (
	// ErrNotSafeConnection the error will force the connection to be closed.
	ErrNotSafeConnection = errors.New("not safe connection")
)

const (
	ctx_conn   = "conn"
	ctx_req_id = "request_id"
	ctx_secret = "secret"
)

type (
	callConnectFunc      func(context.Context, *packet.Msg) (*packet.Msg, error)
	callCloseFunc        func(context.Context)
	callNotRouteFunc     func(context.Context, *packet.Msg)
	callDynamicRouteFunc func(context.Context, *packet.Route, *packet.Msg)
)

type Pulse struct {
	Network     string
	Port        int
	readTimeOut time.Duration

	// buffer pool.
	bufferPool *sync.Pool
	// packet object pool.
	packetPool *sync.Pool

	callConnectFunc      callConnectFunc
	callCloseFunc        callCloseFunc
	callNotRouteFunc     callNotRouteFunc
	callDynamicRouteFunc callDynamicRouteFunc
}

// network: tcp、udp
func New(network string, port int) *Pulse {
	return newPulse(network, port)
}

func newPulse(network string, port int) *Pulse {
	if log.L() == nil {
		log.NewSimple()
	}
	pulse := &Pulse{
		Network: network,
		Port:    port,
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

func (pl *Pulse) ReadTimeOut(t time.Duration) *Pulse {
	pl.readTimeOut = t
	return pl
}

func (pl *Pulse) Listen() error {
	if pl.readTimeOut == 0 {
		pl.readTimeOut = 60 * time.Second
	}
	switch pl.Network {
	case "tcp", "tcp4", "tcp6":
		return pl.listen()
	case "udp", "udp4", "udp6":
		// TODO
		return nil
	default:
		return errors.New("Unsupported network protocol: " + pl.Network)
	}
}

func (pl *Pulse) listen() error {
	addr := fmt.Sprintf(":%d", pl.Port)
	ln, err := net.Listen(pl.Network, addr)
	if err != nil {
		return err
	}

	ipAddr, err := ip.GetLocalIP()
	if err != nil {
		return err
	}
	localAddr := fmt.Sprintf("%s:%d", ipAddr, pl.Port)

	log.L().Info("server started successfully", zap.String("listen", localAddr), zap.String("network", pl.Network))

	pl.accept(ln)
	return nil
}

// receive tcp connection and message.
func (s *Pulse) accept(ln net.Listener) {
	for {
		tcpConn, err := ln.Accept()
		if err != nil {
			log.L().Error(err.Error())
		}

		go s.handle(tcpConn)
	}
}

// handling connections and messages。
func (pl *Pulse) handle(netconn net.Conn) {
	reader := bufio.NewReader(netconn)
	buf := pl.bufferPool.Get().(*bytes.Buffer)
	var (
		offset    int
		varintLen int
		size      int

		ctx = context.Background()
	)
	for {
		// set timeout.
		netconn.SetReadDeadline(time.Now().Add(pl.readTimeOut))
		for {
			b, err := reader.ReadByte()
			if err != nil {
				log.L().Warn(err.Error())
				netconn.Close()

				ctxConn := CtxConn(ctx)
				if ctxConn == nil {
					return
				}

				if mapConn := _connCache.Get(ctxConn.udid); mapConn == nil {
					return
				} else {
					if ctxConn.connectTime < mapConn.connectTime {
						return
					}
					_connCache.Del(ctxConn.udid)
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

				// if len(p.RequestId) == 0 {
				// 	ctx = pl.setCtxReqId(ctx, uuid.New().String())
				// } else {
				// 	ctx = pl.setCtxReqId(ctx, p.RequestId)
				// }

				if p.Type == packet.Type_Connect {
					// type connect,
					// connect type is handled separately.
					ctx = pl.connect(ctx, netconn, p)

					if p.AuthMode == packet.AuthMode_HmacSha1 {
						ctx = pl.setSecret(ctx, p.Secret)
					}

					// server to client.
					pAck := new(packet.Packet)
					pAck.Type = packet.Type_ConnAck
					pAck.Udid = p.Udid
					if pl.callConnectFunc != nil {
						repMsg, err := pl.callConnectFunc(ctx, p.Msg)
						if err == ErrNotSafeConnection {
							log.L().Warn(ErrNotSafeConnection.Error(), zap.String("remote_addr", netconn.RemoteAddr().String()), zap.String("network", netconn.RemoteAddr().Network()))
							netconn.Close()
							return
						}
						if repMsg != nil {
							// type connack,
							// connack type is handled separately.
							pAck.Msg = repMsg
						}

					}
					rbyte, err := Encode(pAck)
					if err != nil {
						log.L().Error(err.Error())
						netconn.Close()
						return
					}
					_, err = netconn.Write(rbyte)
					if err != nil {
						log.L().Error(err.Error())
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

func (pl *Pulse) connect(ctx context.Context, netconn net.Conn, p *packet.Packet) context.Context {
	c := newConn(netconn)
	if len(p.Udid) == 0 {
		c.udid = uuid.New().String()
		p.Udid = c.udid
	} else {
		if conn := _connCache.Get(p.Udid); conn != nil {
			_connCache.Del(p.Udid)
			// if pl.callCloseFunc != nil {
			// 	oldctx := pl.setCtxConn(context.Background(), conn)
			// 	pl.callCloseFunc(oldctx)
			// }
		}
		c.udid = p.Udid
	}
	c.localAddr = p.LocalAddr
	c.network = netconn.RemoteAddr().Network()
	c.remoteAddr = netconn.RemoteAddr().String()
	c.connectTime = time.Now().UnixNano() / 1e6
	_connCache.Put(c.udid, c)
	ctx = pl.setCtxConn(ctx, c)
	return ctx
}

func (pl *Pulse) parse(ctx context.Context, netconn net.Conn, p *packet.Packet) {
	switch p.Type {
	case packet.Type_Ping:
		pl.pong(netconn)
	case packet.Type_RouteMsg:
		pl.route(ctx, p)
	}
}

func (pl *Pulse) pong(netconn net.Conn) {
	wp := new(packet.Packet)
	wp.Type = packet.Type_Pong
	b, err := Encode(wp)
	if err != nil {
		log.L().Error(err.Error())
		return
	}
	netconn.Write(b)
}

func (pl *Pulse) route(ctx context.Context, p *packet.Packet) {
	if p.RouteMode == packet.RouteMode_Not {
		if pl.callNotRouteFunc != nil {
			pl.callNotRouteFunc(ctx, p.Msg)
		}
		return
	}
	if p.RouteMode == packet.RouteMode_Dynamic {
		if pl.callDynamicRouteFunc != nil {
			pl.callDynamicRouteFunc(ctx, p.Route, p.Msg)
		}
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
		rf(ctx, p.Msg)
		return
	}

	f, err := route.GetRouteGroup(p.Route.Id, p.Route.Group)
	if err != nil {
		log.L().Error(err.Error())
		return
	}

	f(ctx, p.Msg)
}

// set connection's context.
func (pl *Pulse) setCtxConn(ctx context.Context, c *Conn) context.Context {
	return context.WithValue(ctx, ctx_conn, c)
}

func (pl *Pulse) setCtxReqId(ctx context.Context, reqId string) context.Context {
	return context.WithValue(ctx, ctx_req_id, reqId)
}

func (pl *Pulse) setSecret(ctx context.Context, secret string) context.Context {
	return context.WithValue(ctx, ctx_secret, secret)
}

func (pl *Pulse) CallConnect(f callConnectFunc) {
	pl.callConnectFunc = f
}

func (pl *Pulse) CallClose(f callCloseFunc) {
	pl.callCloseFunc = f
}

func (pl *Pulse) CallNotRoute(f callNotRouteFunc) {
	pl.callNotRouteFunc = f
}

func (pl *Pulse) CallDynamicRoute(f callDynamicRouteFunc) {
	pl.callDynamicRouteFunc = f
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

func CtxConn(ctx context.Context) *Conn {
	if c := ctx.Value(ctx_conn); c == nil {
		return nil
	} else {
		return c.(*Conn)
	}
}

func CtxRequestId(ctx context.Context) string {
	if c := ctx.Value(ctx_req_id); c == nil {
		return ""
	} else {
		return c.(string)
	}
}

func CtxSecret(ctx context.Context) string {
	if c := ctx.Value(ctx_secret); c == nil {
		return ""
	} else {
		return c.(string)
	}
}

func CacheConn(udid string) *Conn {
	return _connCache.Get(udid)
}

func CacheConns() []*Conn {
	return _connCache.List()
}
