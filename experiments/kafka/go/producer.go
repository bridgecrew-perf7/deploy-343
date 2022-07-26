package main

import (
	"flag"
	"fmt"
	"log"
	// "os"
	"strings"

	"github.com/Shopify/sarama"
)

var (
	_Addrs        []string
	_Topic        string
	_KafkaVersion string
)

func init() {
	// _Addrs = []string{"127.0.0.1:9093"}
	_Topic = "test"
	_KafkaVersion = "3.2.0"
}

func main() {
	var (
		err   error
		addrs string
		index int
		num   int

		config   *sarama.Config
		producer sarama.AsyncProducer
	)

	flag.StringVar(&addrs, "addr", "127.0.0.1:9093", "kakfa brokers address seperated by comma")
	flag.IntVar(&index, "index", 0, "first message index")
	flag.IntVar(&num, "num", 10, "number of messages")
	flag.Parse()

	for _, v := range strings.Split(addrs, ",") {
		v = strings.TrimSpace(v)
		if v != "" {
			_Addrs = append(_Addrs, v)
		}
	}
	fmt.Println("~~~ _Addrs:", _Addrs)

	config = sarama.NewConfig()
	if config.Version, err = sarama.ParseKafkaVersion(_KafkaVersion); err != nil {
		log.Fatalln(err)
	}

	if producer, err = sarama.NewAsyncProducer(_Addrs, config); err != nil {
		log.Fatalln(err)
	}

	for i := index; i < index+num; i++ {
		msg := fmt.Sprintf("hello message: %d", i)
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
