# PULSE

May not be the best performance library for TCP/UDP(currently only TCP), but it is simple and easy to use, suitable for IOT and IM, or more.

## TODO 

Compatibility is not guaranteed before version 1.0.0

## Example

### Server

``` golang
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

func dynamic(ctx context.Context, route *packet.Route, msg *packet.Msg) {
	log.L().Debug(string(msg.Body), zap.Reflect("route", route))
}

func not(ctx context.Context, msg *packet.Msg) {
	log.L().Debug(string(msg.Body))
}

func main() {
	route.ID(1, routeTest)

	pl := pulse.New("tcp", 8080)
	pl.CallConnect(connect)
	pl.CallClose(close)
	pl.CallDynamicRoute(dynamic)
	pl.CallNotRoute(not)
	pl.Listen()
}

```

## Client

``` golang
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
		msg.Body = []byte("not route")
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
		msg2.Body = []byte("dynamic")
		err = c.WriteDynamic(1, "group", msg2)
		if err != nil {
			log.L().Error(err.Error())
		}

		time.Sleep(10 * time.Millisecond)
	}
}

```