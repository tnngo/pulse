package main

import (
	"context"
	"time"

	"github.com/tnngo/log"
	"github.com/tnngo/pulse/client"
	"github.com/tnngo/pulse/packet"
)

func connAck(ctx context.Context, msg *packet.Msg) {
	log.L().Debug(string(msg.Body))
}

func main() {
	log.NewSimple()
	c := client.New("tcp", ":8080")
	c.UDID("123")
	c.CallConnAck(connAck)
	c.Dial()

	for {
		err := c.WriteRoute(1, []byte("test"))
		if err != nil {
			log.L().Error(err.Error())
		}

		time.Sleep(10 * time.Millisecond)
	}

	//select {}
}
