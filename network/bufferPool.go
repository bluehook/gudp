// Copyright 2015 The GUDP Authors. All rights reserved.
// HTTPS clone URL: https://github.com/bluehook/gudp.git
// Blog: http://monsterapp.cn
// Email: bluehook@126.com
package network

import (
	"container/list"
)

//#缓存池
// 回收再利用
type BufferPool struct {
	idleBlockList *list.List
}

// 全局帮助函数
func GetPacketBuffer() []byte {
	return GetBufferPool().GetBlock()
}

func RecoverPacketBuffer(b []byte) {
	GetBufferPool().RecoverBlock(b)
}

// 全局只存在一个
var _instanceBufferPool *BufferPool

func GetBufferPool() *BufferPool {
	if _instanceBufferPool == nil {
		_instanceBufferPool = &BufferPool{idleBlockList: list.New()}
	}
	return _instanceBufferPool
}

// 获取一个PacketBufSize大小的数据块
// 没有就新建一个
func (self *BufferPool) GetBlock() []byte {
	if self.idleBlockList.Len() == 0 {
		return make([]byte, PacketBufSize)
	}
	first := self.idleBlockList.Front().Value.([]byte)
	self.idleBlockList.Remove(self.idleBlockList.Front())
	return first
}

// 回收数据块
func (self *BufferPool) RecoverBlock(b []byte) {
	self.idleBlockList.PushBack(b)
}
