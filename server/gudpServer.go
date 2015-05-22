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

//#可信UDP服务端
// GudpServer为上层业务逻辑提供一层透明的网络服务
// GudpServer维护一个连接列表，内部有自己的包格式
// 发包时对上层数据包进行封装，收包时解封上层数据包
type GudpServer struct {
	net       network.Networker
	group     *network.NetGroup //连接列表
	die       chan bool         //loop线程关闭信号
	timeNs    int64             //时间戳(纳秒)
	elapsedNs int64             //2次更新间隔时间
}

//##创建服务器
func NewGudpServer() *GudpServer {

	gudp := &GudpServer{}
	gudp.net = network.NewNetworkUdp()
	gudp.group = network.NewNetGroup()
	gudp.timeNs = time.Now().UnixNano()

	return gudp
}

// GudpServer主循环
func (self *GudpServer) Update() {
	for {
		select {
		case <-self.die:
			log.Println("handler接收线程终止.")
			return
		default:
			/*开始主循环更新*/
			tmpTimeNs := time.Now().UnixNano()
			self.elapsedNs = tmpTimeNs - self.timeNs
			self.timeNs = tmpTimeNs
		}
	}
}

// 启动住循环更新
func (self *GudpServer) OpenUpdate() {
	self.die = make(chan bool)
	go self.Update()
}

// 关闭主循环更新
func (self *GudpServer) CloseUpdate() {
	close(self.die)
}
