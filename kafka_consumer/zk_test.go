package main

import (
	"fmt"
	"log-collect-master-7bb6ebd327960ddf55c93bdb67e527c352a97c6d/src/github.com/samuel/go-zookeeper/zk"
	"time"
)

func getConnect(zkList []string) (conn *zk.Conn) {
	conn,_,err := zk.Connect(zkList, 10*time.Second)
	if err!=nil {
		fmt.Println(err)
	}
	return
}

func test2(){
	zkList := []string{"localhost:2183"}
	conn := getConnect(zkList)
	defer  conn.Close()
	conn.Create("test",nil,zk.FlagEphemeral,zk.WorldACL(zk.PermAll))
	time.Sleep(20*time.Second)
}

func test3() {
	zkList := []string{"localhost:2183"}
	conn := getConnect(zkList)
	defer conn.Close()
	conn.Create("/test",nil,zk.FlagEphemeral,zk.WorldACL(zk.PermAll))
	time.Sleep(10*time.Second)
}

func main() {
	test3()
}