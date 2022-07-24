package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	var (
		addrs      []string
		topics     []string
		partitions []int32
		err        error
		cancel     chan struct{}

		config    *sarama.Config
		consumer  sarama.Consumer
		pconsumer sarama.PartitionConsumer
	)

	config = sarama.NewConfig()
	addrs = []string{"127.0.0.1:9093"}
	cancel = make(chan struct{})

	if consumer, err = sarama.NewConsumer(addrs, config); err != nil {
		log.Fatalln(err)
	}

	if topics, err = consumer.Topics(); err != nil {
		log.Fatalln(err)
	}

	if len(topics) == 0 {
		log.Fatalln("no topics")
	}
	fmt.Println("~~~ topics:", topics)

	if partitions, err = consumer.Partitions(topics[0]); err != nil {
		log.Fatalln(err)
	}
	if len(partitions) == 0 {
		log.Fatalf("topics %s has no partitions\n", topics[0])
	}
	fmt.Printf("~~~ partitions of %s: %v\n", topics[0], partitions)

	if pconsumer, err = consumer.ConsumePartition("test", 0, 0); err != nil {
		log.Fatalln(err)
	}

	go func() {
		mc := pconsumer.Messages() // *sarama.ConsumerMessage

		for {
			select {
			case msg := <-mc:
				fmt.Printf(
					">>> msg.Timestamp=%+v, msg.Topic=%v, msg.Partition=%v, msg.Offset=%v\n",
					msg.Timestamp, msg.Topic, msg.Partition, msg.Offset,
				)
				fmt.Printf("    key: %q, value: %q\n", msg.Key, msg.Value)
			case <-cancel:
				return
			}
		}
	}()

	time.Sleep(30 * time.Second)
	close(cancel)

	if err = consumer.Close(); err != nil {
		log.Fatalln(err)
	}
}
