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
	r := New()
	//tr := new(TestRoute)
	r.Put(1, nil)
	err := r.Put(1, nil)
	if err != nil {
		t.Log(err)
	}
}
