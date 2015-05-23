// Copyright 2015 The GUDP Authors. All rights reserved.
// HTTPS clone URL: https://github.com/bluehook/gudp.git
// Blog: http://monsterapp.cn
// Email: bluehook@126.com

package main

import (
	"./network"
	"./server"
	//"log"
	"time"
)

//实现工作者接口
type MyWork struct {
}

func (self *MyWork) Update(elapsed int64) bool {
	//log.Println("MyWork Update()")
	return true
}

func (self *MyWork) HandlePacket(packet *network.Packet) {

}

func main() {

	server := server.NewGudpServer()
	server.GudpWorker = &MyWork{}
	server.Start()
	<-time.After(1e9)
	server.Close()
	<-time.After(1e9)
}
