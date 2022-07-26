package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

// TODO: ack

var (
	_KafkaVersion string
	_Addrs        []string
	_GroupId      string
	_Topic        string
)

func init() {
	_Addrs = []string{"127.0.0.1:9093"}
	_GroupId = "default"
	_Topic = "test"
	_KafkaVersion = "3.2.0"
}

func main() {
	var (
		err error
		ctx context.Context

		config  *sarama.Config
		group   sarama.ConsumerGroup
		handler *CGHandler // sarama.ConsumerGroupHandler
	)

	if len(os.Args) > 1 {
		_Addrs = os.Args[1:]
	}
	fmt.Println("~~~ _Addrs:", _Addrs)

	config = sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	if config.Version, err = sarama.ParseKafkaVersion(_KafkaVersion); err != nil {
		log.Fatalln(err)
	}
	group, err = sarama.NewConsumerGroup(_Addrs, _GroupId, config)
	if err != nil {
		log.Fatalln(err)
	}

	ctx = context.Background()
	handler = NewCGHandler(ctx, group)
	handler.Consume(_Topic)

	time.Sleep(15 * time.Second)
	log.Println("<<< Exit")

	handler.Close()
	if err = group.Close(); err != nil {
		log.Fatalln(err)
	}
	time.Sleep(2 * time.Second)
}

type CGHandler struct {
	group  sarama.ConsumerGroup
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
}

func NewCGHandler(ctx context.Context, group sarama.ConsumerGroup) (handler *CGHandler) {
	handler = new(CGHandler)
	handler.group = group
	handler.ctx, handler.cancel = context.WithCancel(ctx)
	handler.wg = new(sync.WaitGroup)

	return handler
}

func (handler *CGHandler) Consume(topics ...string) {
	go func() {
		var err error
		log.Println("==> Handler.Consume start")
		for {
			err = handler.group.Consume(handler.ctx, topics, handler)
			if err != nil {
				if errors.Is(sarama.ErrClosedConsumerGroup, err) {
					log.Println("!!! Handler.Consume closed:", err)
					return
				}
				log.Println("!!! Handler.Consume error:", err)
			} else {
				log.Println("<== Handler.Consume end")
			}

			if handler.ctx.Err() != nil {
				return
			}
		}
	}()
}

func (handler *CGHandler) Close() {
	handler.cancel()
	handler.wg.Wait()
}

func (handler *CGHandler) Setup(sess sarama.ConsumerGroupSession) (err error) {
	log.Println(">>> Handler.Setup Start")

	go func() {
		var err error
		for {
			select {
			case err = <-handler.group.Errors():
				if errors.Is(sarama.ErrClosedConsumerGroup, err) {
					log.Println("!!! Handle.Setup closed")
					return
				}
				log.Println("!!! Handle.Setup error:", err)
			case <-handler.ctx.Done():
				log.Println("~~~ Handle Setup done")
				return
			}
		}
	}()

	return nil
}

func (handler *CGHandler) Cleanup(sess sarama.ConsumerGroupSession) (err error) {
	// TODO
	log.Println(">>> Handler Cleanup")
	return nil
}

func (handler *CGHandler) ConsumeClaim(sess sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim) (err error) {

	handler.wg.Add(1)
	defer handler.wg.Done()

	tmpl := "<-- msg.Timestamp=%+v, msg.Topic=%v, msg.Partition=%v, msg.Offset=%v\n" +
		"    key: %q, value: %q\n"

LOOP:
	for {
		select {
		case msg := <-claim.Messages():
			if msg == nil {
				break LOOP
			}
			log.Printf(
				tmpl, msg.Timestamp, msg.Topic, msg.Partition, msg.Offset,
				msg.Key, msg.Value,
			)

			// TODO: process(msg)
			// sess.MarkOffset(msg.Topic, msg.Partition, msg.Offset, "some-metadata")
			sess.MarkMessage(msg, "consumed-by-d2jvkpn")
		case <-handler.ctx.Done():
			log.Println("!!! ConsumeClaim canceled")
			break LOOP
		}
	}

	return nil
}
