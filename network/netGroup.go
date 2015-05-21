// Copyright 2015 The GUDP Authors. All rights reserved.
// HTTPS clone URL: https://github.com/bluehook/gudp.git
// Blog: http://monsterapp.cn
// Email: bluehook@126.com

package network

import (
	"sync"
)

//#连接对象容器
// 内部包含一个map，带读写锁
type NetGroup struct {
	connGroup map[ConnId]NetConnectioner
	mutex     sync.RWMutex
	base      uint64 //session生成基数
}

//##全局只存在一个连接管理
var _instanceNetGroup *NetGroup

func GetConnGroup() *NetGroup {
	if _instanceNetGroup == nil {
		_instanceNetGroup = &NetGroup{connGroup: make(map[ConnId]NetConnectioner)}
	}
	return _instanceNetGroup
}

//##基本操作
func (self *NetGroup) Add(conn NetConnectioner) {

	self.mutex.Lock()
	defer self.mutex.Unlock()
	self.base++
	id := self.base
	self.connGroup[ConnId(id)] = conn
	//Session高48位为ConnId
	conn.SetSession(ComposeSession(ConnId(id), SessionFlag(SessionFlag_Init)))
}

func (self *NetGroup) Del(id ConnId) {

	self.mutex.Lock()
	defer self.mutex.Unlock()
	delete(self.connGroup, id)
}

func (self *NetGroup) Get(id ConnId) (conn NetConnectioner, ok bool) {
	conn, ok = self.connGroup[id]
	return
}

func (self *NetGroup) GetBySession(session Session) (conn NetConnectioner, ok bool) {
	id := SessionToConnId(session)
	conn, ok = self.connGroup[id]
	return
}
