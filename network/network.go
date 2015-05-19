package network

import (
	"fmt"
	"net"
)

//#UDP框架层次结构
// 1. UDP底层
// 2. 链路层
// 3. 流量控制层
// 4. 数据包处理层

//##UDP底层

// 网络底层接口
type Networker interface {
	Open(port int) bool
	Connect(ip string, port int) bool
	Close()
	GetReadChan() chan *NetworkPacket
	GetWriteChan() chan *NetworkPacket
}

// 底层数据包
type NetworkPacket struct {
	Addr *net.UDPAddr
	Buf  []byte
	Size int
}

// 封装net库的UDP处理
// 实现Networker接口
type NetworkUdp struct {
	conn      *net.UDPConn
	readchan  chan *NetworkPacket
	writechan chan *NetworkPacket
}

// 作为服务端打开监听端口
func (self *NetworkUdp) Open(port int) bool {

	sport := fmt.Sprintf(":%d", port)
	udpAddr, err := net.ResolveUDPAddr("udp4", sport)
	if err == nil {
		self.conn, err = net.ListenUDP("udp", udpAddr)
		if err == nil {
			self.conn.SetReadBuffer(32768)
			self.conn.SetWriteBuffer(32768)
			self.readchan = make(chan *NetworkPacket, 1024)
			self.writechan = make(chan *NetworkPacket, 1024)
			self.handler()
			log("udp server opened.")
			return true
		}
	}
	checkError(err)
	return false
}

// 作为客户端连接目标端口
func (self *NetworkUdp) Connect(ip string, port int) bool {

	addr := fmt.Sprintf("%s:%d", ip, port)
	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err == nil {
		self.conn, err = net.DialUDP("udp", nil, udpAddr)
		if err == nil {
			self.readchan = make(chan *NetworkPacket, 32)
			self.writechan = make(chan *NetworkPacket, 32)
			self.handler()
			log("udp client ready.")
			return true
		}
	}
	checkError(err)
	return false
}

// 关闭底层UDP连接
func (self *NetworkUdp) Close() {
	if self.conn != nil {
		self.conn.Close()
	}
}

// 获取读通道
func (self *NetworkUdp) GetReadChan() chan *NetworkPacket {
	return self.readchan
}

// 获取写通道
func (self *NetworkUdp) GetWriteChan() chan *NetworkPacket {
	return self.writechan
}

// 处理数据
func (self *NetworkUdp) handler() {

	//接收
	go func(net *NetworkUdp) {
		for {
			var buf [512]byte
			num, addr, err := net.conn.ReadFromUDP(buf[0:])
			if err != nil {
				checkError(err)
				return
			}
			net.readchan <- &NetworkPacket{addr, buf[0:], num}
		}
	}(self)

	//发送
	go func(net *NetworkUdp) {
		for {
			packet := <-net.writechan
			if packet.Addr != nil {
				_, err := net.writeto(packet.Buf, packet.Addr)
				if err != nil {
					checkError(err)
				}
			} else {
				_, err := net.write(packet.Buf)
				if err != nil {
					checkError(err)
				}
			}
		}
	}(self)
}

// 发送数据
func (self *NetworkUdp) write(b []byte) (int, error) {
	return self.conn.Write(b)
}

// 发送数据到地址
func (self *NetworkUdp) writeto(b []byte, addr net.Addr) (int, error) {
	return self.conn.WriteTo(b, addr)
}

// 读取数据
func (self *NetworkUdp) read(b []byte) (int, error) {
	return self.conn.Read(b)
}

// 创建UDP网络对象
func NewNetworkUdp() Networker {
	nw := new(NetworkUdp)
	return nw
}

//##链路层

// 三次握手状态定义
const (
	UDPM_INIT = iota //初始化状态
	UDPM_SYN         //同步状态
	UDPM_FN          //完成状态
)

//##流量控制层

//##数据包处理层
