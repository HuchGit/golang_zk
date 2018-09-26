package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
	"log"
	"os"
)

var Address = []string{"bjpg-h287.yz02:9092"}

func main() {
	syncProducer(Address)
}

func syncProducer(address []string){
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5*time.Second
	p ,err := sarama.NewSyncProducer(address,config)
	if err != nil {
		log.Printf("sarama.New err message=%s \n",err)
	}
	defer p.Close()

	topic := "comsumer_perf_log"

	srcValue := "sync: this is a message. index=%d"

	for i:=0;i<10;i++ {
		value := fmt.Sprintf(srcValue,i)
		msg := &sarama.ProducerMessage{
			Topic:topic,
			Value:sarama.ByteEncoder{},
		}
		part,offset,err := p.SendMessage(msg)
		if err != nil {
			log.Printf("send message(%s) err=%s \n", value, err)
		}else {
			fmt.Fprintf(os.Stdout, value + "发送成功，partition=%d, offset=%d \n", part, offset)
		}
		time.Sleep(2*time.Second)
	}
}