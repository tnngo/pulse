package main

import (
	"context"
	"time"

	"github.com/tnngo/log"
	"github.com/tnngo/pulse/client"
)

func connAck(ctx context.Context, body []byte) {
	log.L().Debug(string(body))
}

func main() {
	log.NewSimple()
	c := client.New("tcp", ":8080")
	c.UDID("123")
	c.CallConnAck(connAck)
	go c.Dial()
	time.Sleep(1 * time.Second)
	c.WriteRoute(1, []byte("test"))

	select {}
}
