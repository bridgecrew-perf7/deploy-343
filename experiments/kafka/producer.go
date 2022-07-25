package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

var (
	_Addrs []string
	_Topic string
)

func init() {
	_Addrs = []string{"127.0.0.1:9093"}
	_Topic = "test"
}

func main() {
	var (
		err error

		config   *sarama.Config
		producer sarama.AsyncProducer
	)

	if len(os.Args) > 1 {
		_Addrs = os.Args[1:]
	}
	fmt.Println("~~~", _Addrs)

	config = sarama.NewConfig()

	if producer, err = sarama.NewAsyncProducer(_Addrs, config); err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 3; i++ {
		msg := fmt.Sprintf("hello, world: %d", i)
		log.Println("--> send msg:", msg)

		pmsg := &sarama.ProducerMessage{
			Topic: _Topic,
			Key:   sarama.StringEncoder("e0001"),
			Value: sarama.ByteEncoder([]byte(msg)),
		}

		producer.Input() <- pmsg
	}

	if err = producer.Close(); err != nil {
		log.Fatalln(err)
	}
}
