package main

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/tnngo/log"
	"github.com/tnngo/pulse/client"
	"github.com/tnngo/pulse/packet"
)

func connAck(ctx context.Context, msg *packet.Msg) {
	if msg == nil {
		return
	}
	log.L().Debug(string(msg.Body))
}

func main() {
	log.NewSimple()
	c := client.New("tcp", ":8080")
	c.UDID("123")
	c.CallConnAck(connAck)
	c.Dial()

	for {
		msg := new(packet.Msg)
		msg.RequestId = uuid.New().String()
		msg.Body = []byte("routes")
		err := c.WriteRoute(1, msg)
		if err != nil {
			log.L().Error(err.Error())
		}

		msg1 := new(packet.Msg)
		msg1.RequestId = uuid.New().String()
		msg1.Body = []byte("normal")
		err = c.Write(msg1)
		if err != nil {
			log.L().Error(err.Error())
		}

		msg2 := new(packet.Msg)
		msg2.RequestId = uuid.New().String()
		msg2.Body = []byte("self-processing")
		err = c.WriteSelfProc(1, "group", msg2)
		if err != nil {
			log.L().Error(err.Error())
		}

		time.Sleep(10 * time.Millisecond)
	}

	//select {}
}
