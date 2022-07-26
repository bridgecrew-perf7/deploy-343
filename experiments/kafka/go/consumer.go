package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

var (
	_Addrs []string
)

func init() {
	_Addrs = []string{"127.0.0.1:9093"}
}

func main() {
	var (
		topics     []string
		partitions []int32
		err        error
		cancel     chan struct{}

		config    *sarama.Config
		consumer  sarama.Consumer
		pconsumer sarama.PartitionConsumer
	)

	if len(os.Args) > 1 {
		_Addrs = os.Args[1:]
	}
	fmt.Println("~~~ _Addrs:", _Addrs)

	config = sarama.NewConfig()
	cancel = make(chan struct{})

	if consumer, err = sarama.NewConsumer(_Addrs, config); err != nil {
		log.Fatalln(err)
	}

	if topics, err = consumer.Topics(); err != nil {
		log.Fatalln(err)
	}

	if len(topics) == 0 {
		log.Fatalln("no topics")
	}
	log.Println("~~~ topics:", topics)

	if partitions, err = consumer.Partitions(topics[0]); err != nil {
		log.Fatalln(err)
	}
	if len(partitions) == 0 {
		log.Fatalf("topics %s has no partitions\n", topics[0])
	}
	log.Printf("~~~ partitions of %s: %v\n", topics[0], partitions)

	if pconsumer, err = consumer.ConsumePartition("test", 0, 0); err != nil {
		log.Fatalln(err)
	}

	go func() {
		mc := pconsumer.Messages() // *sarama.ConsumerMessage

		tmpl := "<-- msg.Timestamp=%+v, msg.Topic=%v, msg.Partition=%v, msg.Offset=%v\n" +
			"    key: %q, value: %q\n"

		for {
			select {
			case msg := <-mc:
				log.Printf(
					tmpl, msg.Timestamp, msg.Topic, msg.Partition, msg.Offset,
					msg.Key, msg.Value,
				)
			case <-cancel:
				return
			}
		}
	}()

	time.Sleep(15 * time.Second)
	close(cancel)

	if err = consumer.Close(); err != nil {
		log.Fatalln(err)
	}
}
