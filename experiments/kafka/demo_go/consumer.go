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

		config    *sarama.Config
		consumer  sarama.Consumer
		pconsumer sarama.PartitionConsumer
	)

	config = sarama.NewConfig()
	addrs = []string{"127.0.0.1:9093"}

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

	if pconsumer, err = consumer.ConsumePartition("test", 0, 0); err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 3; i++ {
		msg := <-pconsumer.Messages() // *sarama.ConsumerMessage
		fmt.Printf(">>> msg: %+v\n", msg)
		fmt.Printf("    key: %s, %s\n", msg.Key, msg.Value)
	}

	if err = consumer.Close(); err != nil {
		log.Fatalln(err)
	}
}
