// Copyright 2015 The GUDP Authors. All rights reserved.
// HTTPS clone URL: https://github.com/bluehook/gudp.git
// Blog: http://monsterapp.cn
// Email: bluehook@126.com

package network

import (
	//"net"
	//"io"
	"bytes"
)

//##链路层
// 负责建立连接，同步，超时重建

// 数据包缓存尺寸
const (
	WindowSize = 32
)

// 会话ID: Session = ConnId(48bit) + SessionFlag(16bit)
type Session uint64
type SessionFlag uint16
type ConnId uint64

// Session状态
const (
	SessionFlag_Init = iota
)

func ComposeSession(id ConnId, flag SessionFlag) Session {
	return Session(id<<16) | (Session(flag) & 0xffff)
}

func SessionToConnId(session Session) ConnId {
	return ConnId(session >> 16)
}

func SessionToFlag(session Session) SessionFlag {
	return SessionFlag(session & 0xffff)
}

func (self Session) Equal(other Session) bool {
	myId := SessionToConnId(self)
	otherId := SessionToConnId(other)
	return myId == otherId
}

//###通用更新者接口
type TimeUpdater interface {
	Update(elapsed int64) bool
}

//###连接对象接口
// 服务端客户端通用
// 服务端: 为每个连接创建一个连接对象与客户端的连接对象通信
// 客户端: 只创建一个连接对象与服务端通信
// 作用: 为上层逻辑提供可靠透明的基于协议数据包的网络底层
type NetConnectioner interface {
	SetSession(Session)
	GetSession() Session
	IsConnected() bool //是否已连接
	SetConnected(bool)
	KeepAlive()              //保持连接,重置超时时间
	CheckTimeout(int64) bool //检查超时
	Ping()                   //联通性检查
	Ack()                    //数据包确认
	ProcessRawPacket([]byte) //处理底层包
	TimeUpdater              //更新
}

//###连接对象
// 实现NetConnectioner接口
type NetConn struct {
	session          Session
	bConnected       bool                      //是否已经连接
	pingSendCount    int                       //收到最后一次回应之后累计发出ping次数
	pingRepeatCount  int                       //容许ping累计次数
	lastPingSendTime int64                     //最后一次ping时间
	totalTime        int64                     //累计时间,作为当前时间使用
	bufList          [WindowSize]*bytes.Buffer //数据包缓存列表
}

func (self *NetConn) SetSession(id Session) {
	self.session = id
}

func (self *NetConn) GetSession() Session {
	return self.session
}

func (self *NetConn) SetConnected(b bool) {
	self.bConnected = b
}

func (self *NetConn) IsConnected() bool {
	return self.bConnected
}

func (self *NetConn) CheckTimeout(time int64) bool {

	if !self.bConnected {
		return false
	}
	/*每6秒ping一次*/
	if time > self.lastPingSendTime+6e9 {
		if self.pingSendCount >= self.pingRepeatCount {
			return true
		}
	}
	self.lastPingSendTime = time
	self.pingSendCount++
	self.Ping()

	return false
}

func (self *NetConn) KeepAlive() {
	self.lastPingSendTime = self.totalTime
	self.pingSendCount = 0
}

func (self *NetConn) Ping() {

}

func (self *NetConn) Ack() {

}

func (self *NetConn) ProcessRawPacket([]byte) {

}

func (self *NetConn) Update(elapsed int64) bool {

	//设置累计时间
	self.totalTime += elapsed

	return true
}

//##流量控制层

//##数据包处理层
