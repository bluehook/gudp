package main

import (
	"./network"
	"fmt"
	"time"
)

//testing
func test_networkudp() {
	server := network.NewNetworkUDP()
	server.Open(4321)
	go func(readchan chan *network.NetworkPacket) {
		pack := <-readchan
		fmt.Println("接收数据2", string(pack.Buf))
	}(server.GetReadChan())

	client := network.NewNetworkUDP()
	client.Connect("localhost", 4321)
	pack := new(network.NetworkPacket)
	pack.Buf = []byte("hello server.")
	client.GetWriteChan() <- pack
}

func main() {
	fmt.Println("GUDP")
	test_networkudp()
	for {
		time.Sleep(1e9)
	}
}
