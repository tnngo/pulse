package route

import (
	"context"
	"errors"
	"fmt"

	"github.com/tnngo/log"
	"github.com/tnngo/pulse/packet"
)

// RouteFunc route function type.
type RouteFunc func(context.Context, *packet.Msg) error

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

// ID default route group name is equal to empty.
func ID(id int32, rf RouteFunc) {
	if _, ok := routeMap["pulse"][id]; ok {
		log.L().Sugar().Errorf("id already exists, id=%d", id)
		return
	}

	if routeMap["pulse"] == nil {
		routeMap["pulse"] = make(map[int32]RouteFunc)
	}

	routeMap["pulse"][id] = rf
}

func RouteGroup(group string) (*route, error) {
	if group == "pulse" {
		return nil, errors.New("pulse is default group name")
	}
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

	if routeMap[rg.group] == nil {
		routeMap[rg.group] = make(map[int32]RouteFunc)
	}
	routeMap[rg.group][id] = rf
	return nil
}

func GetRoute(id int32) (RouteFunc, error) {
	if v, ok := routeMap["pulse"][id]; !ok {
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
