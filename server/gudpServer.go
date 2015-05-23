// Copyright 2015 The GUDP Authors. All rights reserved.
// HTTPS clone URL: https://github.com/bluehook/gudp.git
// Blog: http://monsterapp.cn
// Email: bluehook@126.com

package server

import (
	"../network"
	"log"
	"time"
)

//#工作者接口
// 由上层实现，在GudpServer中即插即用
type GudpWorker interface {
	//更新者，返回false表示服务结束,将关闭GudpServer
	network.TimeUpdater
	//数据包处理
	HandlePacket(*network.Packet)
}

//#可信UDP服务
// GudpServer为上层业务逻辑提供一层透明的网络服务
// GudpServer维护一个连接列表，内部有自己的包格式
// 发包时对上层数据包进行封装，收包时解封上层数据包
type GudpServer struct {
	net        network.Networker
	group      *network.NetGroup //连接列表
	die        chan bool         //loop线程关闭信号
	timeNs     int64             //时间戳(纳秒)
	elapsedNs  int64             //2次更新间隔时间
	GudpWorker                   //工作者
}

//##创建服务
func NewGudpServer() *GudpServer {

	gudp := &GudpServer{}
	gudp.net = network.NewNetworkUdp()
	gudp.group = network.NewNetGroup()
	gudp.timeNs = time.Now().UnixNano()

	return gudp
}

// GudpServer主循环
func (self *GudpServer) Update() {

	log.Println("GudpServer主循环开启.")
	for {
		select {
		case <-self.die:
			log.Println("GudpServer主循环退出.")
			return
		default:
			/*开始主循环更新*/
			tmpTimeNs := time.Now().UnixNano()
			self.elapsedNs = tmpTimeNs - self.timeNs
			self.timeNs = tmpTimeNs
			// 推进底层更新
			self.group.Iteration(func(conn network.NetConnectioner) {
				conn.Update(self.elapsedNs)
			})
			// 推进工作者更新
			if !self.GudpWorker.Update(self.elapsedNs) {
				self.Close()
				log.Println("GudpServer主循环主动退出.")
				return
			}
		}
	}
}

// GudpServer数据包接收
func (self *GudpServer) HandlePacket() {

	log.Println("GudpServer数据包接收开启.")
	readchan := self.net.GetReadChan()
	for recv := range readchan {
		select {
		case <-self.die:
			log.Println("GudpServer数据包接收退出.")
			return
		default:
			// 1. 解压底层数据包
			// 2. 生成Packet交上层
			log.Println(recv)
		}
	}
}

// 开始服务
func (self *GudpServer) Start() {
	self.die = make(chan bool)
	if self.GudpWorker != nil {
		//开启主循环
		go self.Update()
		//开启数据接收线程
		go self.HandlePacket()
	}
}

// 关闭服务
func (self *GudpServer) Close() {
	close(self.die)
	self.net.Close()
}
