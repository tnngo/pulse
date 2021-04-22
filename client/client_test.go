package client

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestDial(t *testing.T) {
	c := New("tcp", "localhost:8080")
	go c.Dial()
	time.Sleep(1 * time.Second)
}

func TestClient_Secret(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "base64",
			args: args{
				key:   "AccessKey",
				value: "AccessSecret",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				udid: uuid.New().String(),
			}
			c.Secret(tt.args.key, tt.args.value)
		})
	}
}
