package route

import (
	"context"
	"testing"
)

type TestRoute struct{}

// 登录方法。
func (r *TestRoute) Sign(ctx context.Context, b []byte) {

}

func Test_route_put(t *testing.T) {
	//tr := new(TestRoute)
	//Put(1, nil)
}
