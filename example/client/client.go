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
