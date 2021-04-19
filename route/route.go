package route

import (
	"context"
	"errors"
	"fmt"
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

	_route *route
)

type route struct {
	routeMap map[int32]RouteFunc
}

func ID(id int32, rf RouteFunc) error {
	if rf == nil {
		return ErrRouteNil
	}

	if _route == nil {
		_route = &route{
			routeMap: make(map[int32]RouteFunc),
		}
	}

	if _, ok := _route.routeMap[id]; ok {
		ErrRouteExist = fmt.Errorf("route already exists. routeid=%d", id)
		return ErrRouteExist
	}
	_route.routeMap[id] = rf
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

type RouteGroup struct {
	name string
}

var routeGroupMap = make(map[string]map[int32]RouteFunc)

func NewGroup(groupName string) (*RouteGroup, error) {
	if _, ok := routeGroupMap[groupName]; ok {
		return nil, errors.New("group name repeated: " + groupName)
	}
	routeGroupMap[groupName] = make(map[int32]RouteFunc)

	return &RouteGroup{
		name: groupName,
	}, nil
}

func (rg *RouteGroup) ID(id int32, rf RouteFunc) error {
	if v, ok := routeGroupMap[rg.name]; ok {
		if _, ok := v[id]; ok {
			ErrRouteExist = fmt.Errorf("route already exists. routeName=%s, routeid=%d", rg.name, id)
		}
		return ErrRouteExist
	}
	return nil
}

func (rg *RouteGroup) Get(id int32, group string) (RouteFunc, error) {
	if len(group) == 0 {
		return nil, errors.New("name not nil")
	}
	if v, ok := routeGroupMap[group]; !ok {
		return nil, fmt.Errorf("%s group name does not exist", group)
	} else {
		if v1, ok := v[id]; !ok {
			return nil, fmt.Errorf("%d route id does not exist", id)
		} else {
			return v1, nil
		}
	}
}
