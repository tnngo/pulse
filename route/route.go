package route

import (
	"context"
	"errors"
	"fmt"
)

// RouteFunc route function type.
type RouteFunc func(context.Context, []byte)

var (
	ErrRouteNil = errors.New("Route type value cannot be nil")
	// ErrRouteExist route already exists.
	ErrRouteExist error
	// ErrRouteNotExist route not found.
	ErrRouteNotExist error
	// ErrRouteType route type is error.
	ErrRouteType error
)

type Route struct {
	routeMap map[int32]RouteFunc
}

// New create new route.
func New() *Route {
	return &Route{
		routeMap: make(map[int32]RouteFunc),
	}
}

// Put
func (r *Route) Put(key int32, rf RouteFunc) error {
	if rf == nil {
		return ErrRouteNil
	}

	if _, ok := r.routeMap[key]; ok {
		ErrRouteExist = fmt.Errorf("route already exists. routeid=%d", key)
		return ErrRouteExist
	}
	r.routeMap[key] = rf
	return nil
}

// Get
func (r *Route) Get(key int32) (RouteFunc, error) {
	if v, ok := r.routeMap[key]; !ok {
		ErrRouteNotExist = fmt.Errorf("route not found：%d", key)
		return nil, ErrRouteExist
	} else {
		return v, nil
	}
}
