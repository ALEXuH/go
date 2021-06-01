package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 3

	client, err := sarama.NewSyncProducer([]string{"10.50.2.56:19092", "10.50.2.56:19091", "10.50.2.56:19093"}, config)

	sarama.NewAsyncProducer()
	if err != nil {
		fmt.Println(" Producer err ", err.Error())
		return
	}
	fmt.Println("dsa")
	defer client.Close()

	// 构造消息
	//msg := &sarama.ProducerMessage{}
	//msg.Topic = "web_log"
	//msg.Value = sarama.StringEncoder(" this is message 111")
	client.SendMessage(&sarama.ProducerMessage{
		Topic: "gotest",
		Value: sarama.StringEncoder("this is message "),
	})

}
