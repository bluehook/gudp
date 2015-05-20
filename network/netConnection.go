package network

import (
	"net"
)

//##链路层
// 基于滑动窗口的可信UDP
// 负责建立连接，同步，超时重建

// Session = ConnId(48bit) + SessionFlag(16bit)
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

// 连接对象接口
type NetConnectioner interface {
	SetSession(id Session)
	GetSession() Session
	GetUdpAddr() *net.UDPAddr
}

//##流量控制层

//##数据包处理层
