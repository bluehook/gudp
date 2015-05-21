package main

import (
	"./network"
	"container/list"
	"fmt"
	"strconv"
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

	fmt.Println("GUDP")
	test_networkudp()
	for {
		<-time.After(1e9)
	}
}
