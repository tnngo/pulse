# PULSE

May not be the best performance library for TCP/UDP(currently only TCP), but it is simple and easy to use, suitable for IOT and IM, or more.

## TODO 

Compatibility is not guaranteed before version 1.0.0

## Example

### Server

``` golang
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
	route.ID(1, routeTest)

	pl := pulse.New("tcp", 8080)
	pl.CallConnect(connect)
	pl.CallClose(close)
	pl.Listen()
}
```

## Client

``` golang
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

```