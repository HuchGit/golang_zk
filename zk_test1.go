package main

import (
    "fmt"
    "github.com/samuel/go-zookeeper/zk"
//    "strings"
    "time"
)

func must(err error) {
    if err != nil {
        panic(err)
    }
}

func connect() *zk.Conn {
    zkList := []string{"zk.test"}
    conn, _, err := zk.Connect(zkList, time.Second)
    must(err)
    return conn
}

func main() {
    conn := connect()
    defer conn.Close()

    // create
    //flags := int32(0)
    //acl := zk.WorldACL(zk.PermAll)

    //path, err := conn.Create("/01", []byte("data"), flags, acl)
    //must(err)

    //fmt.Printf("create: %+v\n", path)

    // get
    data, stat, err := conn.Get("/kuaishou/kafkaLogs/producers/PG_LOG2/perf_log")
    must(err)
    fmt.Printf("get: %+v %+v\n", string(data), stat)

    // set
    str := "name=test\npasswad=123456"
    stat, err = conn.Set("/kuaishou/kafkaLogs/producers/PG_LOG2/perf_log", []byte(str), stat.Version)
    must(err)
    fmt.Printf("set: %+v\n", stat)

    // get
    data, stat, err = conn.Get("/kuaishou/kafkaLogs/producers/PG_LOG2/perf_log")
    must(err)
    fmt.Printf("get: %+v %+v\n", string(data), stat)

    // delete
//    err = conn.Delete("/kuaishou/kafkaLogs/producers/PG_LOG2/perf_log", -1)
//    must(err)
//    fmt.Printf("delete: ok\n")

    // exists
//    exists, stat, err := conn.Exists("/kuaishou/kafkaLogs/producers/PG_LOG2/perf_log")
//    must(err)
//    fmt.Printf("exists: %+v %+v\n", exists, stat)
}
