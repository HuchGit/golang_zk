package main

import (
        "fmt"
        zk "github.com/samuel/go-zookeeper/zk"
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
        zkList := []string{"test"}
        conn := getConnect(zkList)
        defer  conn.Close()
        conn.Create("test",nil,zk.FlagEphemeral,zk.WorldACL(zk.PermAll))
        time.Sleep(20*time.Second)
}

func test3() {
        zkList := []string{"zk.cluster.yz"}
        conn := getConnect(zkList)
        defer conn.Close()
        contents, _, ch, err := conn.GetW("/kafkaLogs/producers/LOG2")
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(string(contents))
        fmt.Println(ch)
        time.Sleep(10*time.Second)
}

func main() {
        test3()
}
