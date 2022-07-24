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
	fmt.Println("~~~ topics:", topics)

	if partitions, err = consumer.Partitions("test"); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("~~~ partitions of test:", partitions)

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
