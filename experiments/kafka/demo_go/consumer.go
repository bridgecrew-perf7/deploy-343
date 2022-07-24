package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	var (
		addrs      []string
		topics     []string
		partitions []int32
		err        error

		config   *sarama.Config
		consumer sarama.Consumer
	)

	config = sarama.NewConfig()
	addrs = []string{"127.0.0.1:9092"}

	if consumer, err = sarama.NewConsumer(addrs, config); err != nil {
		log.Fatalln(err)
	}

	if topics, err = consumer.Topics(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("~~~ topics:", topics)

	if partitions, err = consumer.Partitions("test"); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("~~~ partitions of test:", partitions)

	if err = consumer.Close(); err != nil {
		log.Fatalln(err)
	}
}
