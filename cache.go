package pulse

import (
	"sync"
)

type cache struct {
	connMap map[string]*Conn
	rwmutex sync.RWMutex
}

var (
	_connCache = &cache{connMap: make(map[string]*Conn)}
)

func (cp *cache) Put(udid string, c *Conn) {
	cp.rwmutex.Lock()
	defer cp.rwmutex.Unlock()
	cp.connMap[udid] = c
}

func (cp *cache) Del(udid string) {
	cp.rwmutex.Lock()
	defer cp.rwmutex.Unlock()
	if c, ok := cp.connMap[udid]; ok {
		c.netconn.Close()
	}
	delete(cp.connMap, udid)
}

func (cp *cache) Get(udid string) *Conn {
	cp.rwmutex.RLock()
	defer cp.rwmutex.RUnlock()
	return cp.connMap[udid]
}

func (cp *cache) List() []*Conn {
	cp.rwmutex.RLock()
	defer cp.rwmutex.RUnlock()
	conns := make([]*Conn, 0)
	for _, v := range cp.connMap {
		conns = append(conns, v)
	}
	return conns
}
