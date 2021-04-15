package route

import (
	"context"
	"errors"
	"fmt"
)

// RouteFunc route function type.
type RouteFunc func(context.Context, []byte) ([]byte, error)

var (
	ErrRouteNil = errors.New("Route type value cannot be nil")
	// ErrRouteExist route already exists.
	ErrRouteExist error
	// ErrRouteNotExist route not found.
	ErrRouteNotExist error
	// ErrRouteType route type is error.
	ErrRouteType error

	_route *route
)

type route struct {
	routeMap map[int32]RouteFunc
}

func Put(key int32, rf RouteFunc) error {
	if rf == nil {
		return ErrRouteNil
	}

	if _route == nil {
		_route = &route{
			routeMap: make(map[int32]RouteFunc),
		}
	}

	if _, ok := _route.routeMap[key]; ok {
		ErrRouteExist = fmt.Errorf("route already exists. routeid=%d", key)
		return ErrRouteExist
	}
	_route.routeMap[key] = rf
	return nil
}

func Get(key int32) (RouteFunc, error) {
	if v, ok := _route.routeMap[key]; !ok {
		ErrRouteNotExist = fmt.Errorf("route not found：%d", key)
		return nil, ErrRouteExist
	} else {
		return v, nil
	}
}
