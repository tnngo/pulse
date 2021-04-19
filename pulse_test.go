package pulse

import (
	"net"
	"testing"
	"time"

	"github.com/tnngo/log"
)

func Test_pulse_Listen(t *testing.T) {
	log.NewSimple()
	p := &Pulse{
		Network: "tcp",
		Port:    8080,
	}
	if err := p.listen(); err != nil {
		t.Error(err)
	}
}

func Test_pulse_handle(t *testing.T) {
	log.NewSimple()
	s, c := net.Pipe()
	go func() {
		s.Close()
	}()
	p := &Pulse{
		Network:     "tcp",
		Port:        8080,
		readTimeOut: 1 * time.Second,
	}
	p.handle(c)
}
