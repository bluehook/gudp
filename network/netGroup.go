package network

import (
	"sync"
)

//#连接对象容器
// 内部包含一个map，带读写锁
type NetGroup struct {
	sessionGroup map[Session]NetConnectioner
	mutex        sync.RWMutex
	base         int64 //session生成基数
}

//##单例
var _instanceNetGroup *NetGroup

func GetConnGroup() *NetGroup {
	if _instanceNetGroup == nil {
		_instanceNetGroup = &NetGroup{sessionGroup: make(map[Session]NetConnectioner)}
	}
	return _instanceNetGroup
}

//##基本操作
func (self *NetGroup) Add(conn NetConnectioner) {

	self.mutex.Lock()
	defer self.mutex.Unlock()
	id := Session(^self.base + 1)
	self.sessionGroup[id] = conn
	conn.SetSession(id)
	self.base++
}

func (self *NetGroup) Del(id Session) {

	self.mutex.Lock()
	defer self.mutex.Unlock()
	delete(self.sessionGroup, id)
}

func (self *NetGroup) Get(id Session) (conn NetConnectioner, ok bool) {
	conn, ok = self.sessionGroup[id]
	return
}
