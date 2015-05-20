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

func testingList() (int64, int64, int64) {

	cur := time.Now().UnixNano()
	testList := list.New()
	for i := 0; i < 1e5; i++ {
		testList.PushBack("id" + strconv.Itoa(i))
	}
	use1 := (time.Now().UnixNano() - cur) / 1e6
	cur = time.Now().UnixNano()
	nameNum := 0
	for e := testList.Front(); e != nil; e = e.Next() {
		nameNum += len(e.Value.(string))
	}
	use2 := (time.Now().UnixNano() - cur) / 1e6
	cur = time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		for e := testList.Front(); e != nil; e = e.Next() {
			if e.Value == "id100" {
				e.Value = "id3333"
			}
		}
	}
	use3 := (time.Now().UnixNano() - cur) / 1e6
	return use1, use2, use3
}

func testingMap() (int64, int64, int64) {

	testMap := make(map[string]string)
	cur := time.Now().UnixNano()

	for i := 0; i < 1e5; i++ {
		idname := "id" + strconv.Itoa(i)
		testMap[idname] = idname
	}
	use1 := (time.Now().UnixNano() - cur) / 1e6
	cur = time.Now().UnixNano()
	nameNum := 0

	for _, name := range testMap {
		nameNum += len(name)
	}
	use2 := (time.Now().UnixNano() - cur) / 1e6
	cur = time.Now().UnixNano()
	for i := 0; i < 1e5; i++ {
		testMap["id43223"] = "id3333"
	}
	use3 := (time.Now().UnixNano() - cur) / 1e6
	return use1, use2, use3
}

func testingListAndMap() {

	var count1, count2, count3 int64
	for i := 0; i < 50; i++ {
		use1, use2, use3 := testingMap()
		count1 += use1
		count2 += use2
		count3 += use3
	}
	count1 /= 50
	count2 /= 50
	count3 /= 50
	fmt.Println("MAP插入100000个元素,遍历,查询分别用时（毫秒）(50次平均值)：", count1, count2, count3)

	count1, count2, count3 = 0, 0, 0
	for i := 0; i < 50; i++ {
		use1, use2, use3 := testingList()
		count1 += use1
		count2 += use2
		count3 += use3
	}
	count1 /= 50
	count2 /= 50
	count3 /= 50
	fmt.Println("List平均插入100000个元素,遍历,查询分别用时（毫秒）(50次平均值)：", count1, count2, ">1分钟")
}

func main() {

	fmt.Println("GUDP")
	//testingListAndMap()
	//test_networkudp()
	for {
		<-time.After(1e9)
	}
}
