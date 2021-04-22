package route

import (
	"context"
	"errors"
	"fmt"

	"github.com/tnngo/log"
)

// RouteFunc route function type.
type RouteFunc func(context.Context, []byte) error

var (
	ErrRouteNil = errors.New("Route type value cannot be nil")
	// ErrRouteExist route already exists.
	ErrRouteExist error
	// ErrRouteNotExist route not found.
	ErrRouteNotExist error
	// ErrRouteType route type is error.
	ErrRouteType error
)

type route struct {
	group string
}

var routeMap = make(map[string]map[int32]RouteFunc)

// ID default ROUTE GROUP NAME is equal to empty.
func ID(id int32, rf RouteFunc) {
	if _, ok := routeMap[""][id]; ok {
		log.L().Sugar().Errorf("id already exists, id=%d", id)
		return
	}

	routeMap[""][id] = rf
}

func RouteGroup(group string) (*route, error) {
	if _, ok := routeMap[group]; ok {
		return nil, errors.New("group name repeated: " + group)
	}
	routeMap[group] = make(map[int32]RouteFunc)

	return &route{
		group: group,
	}, nil
}

func (rg *route) ID(id int32, rf RouteFunc) error {
	if v, ok := routeMap[rg.group]; ok {
		if _, ok := v[id]; ok {
			ErrRouteExist = fmt.Errorf("route already exists. group=%s", rg.group)
		}
		return ErrRouteExist
	}
	return nil
}

func GetRoute(id int32) (RouteFunc, error) {
	if v, ok := routeMap[""][id]; !ok {
		return nil, fmt.Errorf("route %d does not exist", id)
	} else {
		return v, nil
	}
}

func GetRouteGroup(id int32, group string) (RouteFunc, error) {
	if len(group) == 0 {
		return nil, errors.New("group not nil")
	}
	if v, ok := routeMap[group]; !ok {
		return nil, fmt.Errorf("group %s does not exist", group)
	} else {
		if v1, ok := v[id]; !ok {
			return nil, fmt.Errorf("route %d does not exist", id)
		} else {
			return v1, nil
		}
	}
}
