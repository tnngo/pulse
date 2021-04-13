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
	c.EnableRequestId()
	c.CallConnAck(connAck)
	go c.Dial()
	time.Sleep(1 * time.Second)
	c.WriteRoute(1, []byte("test"))

	//p := new(packet.Packet)
	//p.Type = packet.Type_Connack
	//b1, err := proto.Marshal(p)
	//if err != nil {
	//	log.L().Error(err.Error())
	//	return
	//}
	//b2 := proto.EncodeVarint(uint64(len(b1)))
	//b2 = append(b2, b1...)
	//conn.Write(b2)

	// b, err := os.ReadFile("./thumb.jpeg")
	// if err != nil {
	// 	log.L().Error(err.Error())
	// 	return
	// }

	// log.L().Debug("图片大小", zap.Int("b", len(b)))

	// p := new(packet.Packet)
	// p.Msg = b
	// b1, err := proto.Marshal(p)
	// if err != nil {
	// 	log.L().Error(err.Error())
	// 	return
	// }

	// b2 := proto.EncodeVarint(uint64(len(b1)))
	// b2 = append(b2, b1...)
	// log.L().Debug("protobuf大小", zap.Int("b2", len(b2)))
	// for {
	// 	time.Sleep(10 * time.Millisecond)
	// 	conn.Write(b2)
	// }

	select {}
}
