package network

import (
//"net"
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

//###连接对象接口
// 服务端客户端通用
// 服务端: 为每个连接创建一个连接对象与客户端的连接对象通信
// 客户端: 只创建一个连接对象与服务端通信
// 作用: 为上层逻辑提供可靠透明的基于协议数据包的网络底层
type NetConnectioner interface {
	SetSession(id Session)
	GetSession() Session
	Send([]byte)
	Recv([]byte)
	IsConnected() bool //是否已连接
	SetConnected(bool)
	KeepAlive()         //保持连接,重置超时时间
	CheckTimeout()      //检查超时
	Ping()              //联通性检查
	Ack()               //数据包确认
	BuildPacketHeader() //生成包头
}

//##流量控制层

//##数据包处理层
