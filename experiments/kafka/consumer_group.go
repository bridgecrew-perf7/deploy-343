package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

// TODO: ack

var (
	_Addrs   []string
	_GroupId string
	_Topic   string
)

func init() {
	_Addrs = []string{"127.0.0.1:9093"}
	_GroupId = "default"
	_Topic = "test"
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
	fmt.Println("~~~", _Addrs)

	config = sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	group, err = sarama.NewConsumerGroup(_Addrs, _GroupId, config)
	if err != nil {
		log.Fatalln(err)
	}

	ctx = context.Background()
	handler = NewCGHandler(ctx)

	go func() {
		var err error
		log.Println("==> ConsumerGroup A start")
		for {
			err = group.Consume(ctx, []string{_Topic}, handler)
			if err != nil {
				if errors.Is(sarama.ErrClosedConsumerGroup, err) {
					log.Println("!!! ConsumerGroup A closed:", err)
					return
				}
				log.Println("!!! ConsumerGroup A error:", err)
			} else {
				log.Println("<== ConsumerGroup A end")
			}

			if ctx.Err() != nil {
				return
			}
		}
	}()

	go func() {
		var err error
		for {
			select {
			case err = <-group.Errors():
				if errors.Is(sarama.ErrClosedConsumerGroup, err) {
					log.Println("!!! ConsumerGroup B closed")
					return
				}
				log.Println("!!! ConsumerGroup B error:", err)
			case <-handler.ctx.Done():
				log.Println("~~~ handler done")
				return
			}
		}
	}()

	time.Sleep(15 * time.Second)
	log.Println("<<< Exit")

	handler.Close()
	if err = group.Close(); err != nil {
		log.Fatalln(err)
	}
	time.Sleep(2 * time.Second)
}

type CGHandler struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewCGHandler(ctx context.Context) (cgh *CGHandler) {
	cgh = new(CGHandler)
	cgh.ctx, cgh.cancel = context.WithCancel(ctx)

	return cgh
}

func (cgh *CGHandler) Close() {
	cgh.cancel()
}

func (cgh *CGHandler) Setup(sess sarama.ConsumerGroupSession) (err error) {
	// TODO
	log.Println(">>> CGHandler Setup")
	return nil
}

func (cgh *CGHandler) Cleanup(sess sarama.ConsumerGroupSession) (err error) {
	// TODO
	log.Println(">>> CGHandler Cleanup")
	return nil
}

func (cgh *CGHandler) ConsumeClaim(sess sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim) (err error) {

	tmpl := "--> msg.Timestamp=%+v, msg.Topic=%v, msg.Partition=%v, msg.Offset=%v\n" +
		"    key: %q, value: %q\n"

LOOP:
	for {
		select {
		case msg := <-claim.Messages():
			if msg == nil {
				break LOOP
			}
			log.Printf(
				tmpl,
				msg.Timestamp, msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value,
			)
		case <-cgh.ctx.Done():
			log.Println("!!! Consumer canceled")
			break LOOP
		}
	}

	return nil
}
