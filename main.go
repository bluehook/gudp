// Copyright 2015 The GUDP Authors. All rights reserved.
// HTTPS clone URL: https://github.com/bluehook/gudp.git
// Blog: http://monsterapp.cn
// Email: bluehook@126.com

package main

import (
	"./network"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"time"
)

//testing
func test_networkudp() {

	server := network.NewNetworkUdp()
	server.Open(4321)
	die := make(chan bool)
	go func(readchan chan *network.NetworkPacket, d chan bool) {
		for {
			select {
			case <-die:
				fmt.Println("MAIN接受数据线程关闭.")
				return
			case pack := <-readchan:
				fmt.Println("接收数据:", string(pack.Buf))
			}

		}
	}(server.GetReadChan(), die)

	client := network.NewNetworkUdp()
	client.Connect("localhost", 4321)
	pack := new(network.NetworkPacket)
	pack.Buf = []byte("hello server.")
	client.GetWriteChan() <- pack
	client.GetWriteChan() <- pack
	client.GetWriteChan() <- pack
	client.GetWriteChan() <- pack
	client.GetWriteChan() <- pack

	time.Sleep(1e9)
	close(die)
	server.Close()
	client.Close()

	time.Sleep(1e9)
	server.Open(4321)
	die = make(chan bool)
	go func(readchan chan *network.NetworkPacket, d chan bool) {
		for {
			select {
			case <-die:
				fmt.Println("MAIN接受数据线程关闭.")
				return
			case pack := <-readchan:
				fmt.Println("接收数据:", string(pack.Buf))
			}

		}
	}(server.GetReadChan(), die)
	client.Connect("localhost", 4321)
	pack = new(network.NetworkPacket)
	pack.Buf = []byte("hello server2.")
	client.GetWriteChan() <- pack
	client.GetWriteChan() <- pack
	client.GetWriteChan() <- pack
	client.GetWriteChan() <- pack
	client.GetWriteChan() <- pack

	time.Sleep(1e9)
	close(die)
	server.Close()
	client.Close()
}

func main() {

	log.Println("GUDP")
	//test_networkudp()

	packet := network.NewPacket(binary.LittleEndian)
	packet.SetBuf(new(bytes.Buffer))

	for {
		<-time.After(1e9)
	}
}
