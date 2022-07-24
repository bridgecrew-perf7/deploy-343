package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	var (
		addrs []string
		err   error

		config   *sarama.Config
		producer sarama.AsyncProducer
	)

	config = sarama.NewConfig()
	addrs = []string{"127.0.0.1:9093"}

	if producer, err = sarama.NewAsyncProducer(addrs, config); err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 3; i++ {
		msg := fmt.Sprintf("hello, world: %d", i)
		fmt.Println(">>> send msg:", msg)

		pmsg := &sarama.ProducerMessage{
			Topic: "test",
			Key:   sarama.StringEncoder("e0001"),
			Value: sarama.ByteEncoder([]byte(msg)),
		}

		producer.Input() <- pmsg
	}

	if err = producer.Close(); err != nil {
		log.Fatalln(err)
	}
}
