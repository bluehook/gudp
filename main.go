package main

import (
	"./network"
	"fmt"
	"time"
)

//testing
func test_networkudp() {

	server := network.NewNetworkUdp()
	server.Open(4321)
	go func(readchan chan *network.NetworkPacket) {
		for {
			pack := <-readchan
			fmt.Println("接收数据:", string(pack.Buf))
			time.Sleep(1e9)
		}
	}(server.GetReadChan())

	client := network.NewNetworkUdp()
	client.Connect("localhost", 4321)
	pack := new(network.NetworkPacket)
	pack.Buf = []byte("hello server.")
	client.GetWriteChan() <- pack
	client.GetWriteChan() <- pack
	client.GetWriteChan() <- pack
	client.GetWriteChan() <- pack
}

func main() {
	fmt.Println("GUDP")
	test_networkudp()
	for {
		time.Sleep(1e9)
	}
}
