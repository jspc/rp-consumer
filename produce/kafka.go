package main

import (
	"bytes"
	"io"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type kafkaClient interface {
	Produce(*kafka.Message, chan kafka.Event) error
	Events() chan kafka.Event
	Flush(int) int
}

type Kafka struct {
	client kafkaClient
}

type marshaller interface {
	Marshall(io.Writer) error
}

func NewKafka(bootstrapServers string) (k Kafka, err error) {
	k.client, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
	})

	return
}

func (k Kafka) FollowLogs() {
	for e := range k.client.Events() {
		log.Print(e.String())
	}
}

func (k *Kafka) Write(topic string, e marshaller) (err error) {
	b := new(bytes.Buffer)

	err = e.Marshall(b)
	if err != nil {
		return
	}

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          b.Bytes(),
		Headers:        []kafka.Header{},
	}

	err = k.client.Produce(message, nil)

	if kafkaError, ok := err.(kafka.Error); ok && kafkaError.Code() == kafka.ErrQueueFull {
		log.Print("Kafka local queue full error - Going to Flush then retry...")
		flushedMessages := k.client.Flush(30 * 1000)

		log.Print("Flushed kafka messages. Outstanding events still un-flushed: %d", flushedMessages)

		return k.Write(topic, e)
	}

	return err
}

func (k Kafka) logs() {
	for m := range k.client.Events() {
		log.Print(m)
	}
}
