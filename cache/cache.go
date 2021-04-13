package cache

import (
	"sync"

	"github.com/tnngo/pulse/conn"
)

func New() *ConnCache {
	return &ConnCache{
		connMap: make(map[string]*conn.Conn),
	}
}

type ConnCache struct {
	connMap map[string]*conn.Conn
	rwmutex sync.RWMutex
}

func (cp *ConnCache) Put(udid string, c *conn.Conn) {
	cp.rwmutex.Lock()
	defer cp.rwmutex.Unlock()
	cp.connMap[udid] = c
}

func (cp *ConnCache) Del(udid string) {
	cp.rwmutex.Lock()
	defer cp.rwmutex.Unlock()
	delete(cp.connMap, udid)
}

func (cp *ConnCache) Get(udid string) *conn.Conn {
	cp.rwmutex.RLock()
	defer cp.rwmutex.RUnlock()
	return cp.connMap[udid]
}

func (cp *ConnCache) List() []*conn.Conn {
	cp.rwmutex.RLock()
	defer cp.rwmutex.RUnlock()
	conns := make([]*conn.Conn, 0)
	for _, v := range cp.connMap {
		conns = append(conns, v)
	}
	return conns
}
