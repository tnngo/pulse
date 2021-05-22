package main

import (
	"context"

	"github.com/tnngo/log"
	"github.com/tnngo/pulse"
	"github.com/tnngo/pulse/packet"
	"github.com/tnngo/pulse/route"
	"go.uber.org/zap"
)

func connect(ctx context.Context, msg *packet.Msg) (*packet.Msg, error) {
	conn := pulse.CtxConn(ctx)
	log.L().Debug("device online", zap.String("udid", conn.UDID()))
	ack := new(packet.Msg)
	ack.Body = []byte("i am server")
	return ack, nil
}

func close(ctx context.Context) {
	conn := pulse.CtxConn(ctx)
	log.L().Debug("device offline", zap.String("udid", conn.UDID()))
}

func routeTest(ctx context.Context, msg *packet.Msg) error {
	log.L().Debug(string(msg.Body))
	return nil
}

func forward(ctx context.Context, route *packet.Route, msg *packet.Msg) {
	log.L().Debug(string(msg.Body))
}

func main() {
	route.ID(1, routeTest)

	pl := pulse.New("tcp", 8080)
	pl.CallConnect(connect)
	pl.CallClose(close)
	pl.CallForward(forward)
	pl.Listen()
}
