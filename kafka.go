package main

import (
	"bytes"
	"io"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type consumer interface {
	SubscribeTopics([]string, kafka.RebalanceCb) error
	ReadMessage(time.Duration) (*kafka.Message, error)
}

type Message struct {
	topic string
	ct    consumerType
}

type consumerType interface {
	Unmarshall(io.Reader) error
}

type Kafka struct {
	c        Config
	consumer consumer
}

func NewKafka(c Config) (k Kafka, err error) {
	hostname, err := os.Hostname()
	if err != nil {
		return
	}

	k.c = c

	k.consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    c.Broker,
		"group.id":             "ingestor",
		"auto.offset.reset":    "earliest",
		"session.timeout.ms":   30000,
		"enable.partition.eof": true,
		"client.id":            hostname,
	})

	if err != nil {
		return
	}

	err = k.consumer.SubscribeTopics(c.Topics(), nil)

	return
}

func (k Kafka) ConsumerLoop(c chan Message) (err error) {
	var msg *kafka.Message

	for {
		msg, err = k.consumer.ReadMessage(time.Millisecond * 250)
		if err != nil {
			if !err.(kafka.Error).IsTimeout() {
				return
			}

			continue
		}

		if msg == nil || msg.Value == nil {
			continue
		}

		topic := *msg.TopicPartition.Topic

		t, err := k.c.Type(topic)
		if err != nil {
			continue
		}

		b := bytes.NewBuffer(msg.Value)
		err = t.Unmarshall(b)
		if err != nil {
			continue
		}

		c <- Message{
			topic: topic,
			ct:    t,
		}
	}
}
