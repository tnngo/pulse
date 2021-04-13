package client

import (
	"testing"
	"time"
)

func TestDial(t *testing.T) {
	c := New("tcp", "localhost:8080")
	go c.Dial()
	time.Sleep(1 * time.Second)
}
