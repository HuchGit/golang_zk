package main

import (
	"fmt"
	"github.com/bsm/sarama-cluster"
	"os"
	"os/signal"
	"sync"
	"log"
	"time"
)

func main(){
	topic := []string{"comsumer_perf_log"}
	var wg = &sync.WaitGroup{}
	wg.Add(2)
	go clusterconsumer(wg,Address,topic,"group-1")
	go clusterconsumer(wg,Address,topic,"group-2")
	wg.Wait()
}

func clusterconsumer(wg *sync.WaitGroup,brokers,topics []string,groupId string){
	defer wg.Done()
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = 1*time.Second

	consumer,err := cluster.NewConsumer(brokers, groupId,topics,config)

	if err != nil {
		log.Printf("%s,sarama,NewSyncProducer err,message = %s\n",groupId,err)
		return
	}
	defer consumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func(){
		for err := range consumer.Errors(){
			log.Printf("%s:Error: %s\n",groupId,err.Error())
		}
	}()
	var successes int
	Loop:
	for {
		select {
			case msg,ok := <-consumer.Messages():
				if ok{
					fmt.Fprintf(os.Stdout,"%s:%s/%d/%d\t%s\t%s\n", groupId, msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
					consumer.MarkOffset(msg,"")
					successes++
				}
			case <-signals:
					break Loop
		}
	}
	fmt.Fprintf(os.Stdout,"%s consume %d messages \n", groupId, successes)
}