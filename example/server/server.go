package main

import (
	"context"

	"github.com/tnngo/log"
	"github.com/tnngo/pulse"
	"github.com/tnngo/pulse/route"
	"go.uber.org/zap"
)

func connect(ctx context.Context, body []byte) ([]byte, error) {
	conn := pulse.CtxConn(ctx)
	log.L().Debug("device online", zap.String("udid", conn.UDID()))
	return nil, nil
}

func close(ctx context.Context) {
	conn := pulse.CtxConn(ctx)
	log.L().Debug("device offline", zap.String("udid", conn.UDID()))
}

func routeTest(ctx context.Context, body []byte) error {
	log.L().Debug(string(body))
	return nil
}

func main() {
	route.Put(1, routeTest)

	pl := pulse.New("tcp", 8080)
	pl.CallConnect(connect)
	pl.CallClose(close)
	pl.Listen()
}
