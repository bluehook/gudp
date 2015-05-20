package network

import (
	"net"
)

//##链路层
// 基于滑动窗口的可信UDP
// 负责建立连接，同步，超时重建
type Session int64

//###连接对象接口
type NetConnectioner interface {
	SetSession(id Session)
	GetSession() Session
	GetUdpAddr() *net.UDPAddr
}

//##流量控制层

//##数据包处理层
