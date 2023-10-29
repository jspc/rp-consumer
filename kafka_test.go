package main

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/jspc/rp-consumer/types"
)

type dummyKafka struct {
	topic   string
	err     bool
	message []byte
}

func (d dummyKafka) SubscribeTopics([]string, kafka.RebalanceCb) error {
	if d.err {
		return fmt.Errorf("en error")
	}

	return nil
}

func (d dummyKafka) ReadMessage(time.Duration) (*kafka.Message, error) {
	if d.err {
		return nil, fmt.Errorf("an error")
	}

	return &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &d.topic, Partition: kafka.PartitionAny},
		Value:          d.message,
		Headers:        []kafka.Header{},
	}, nil
}

func TestNewKafka(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("unexpected error: %+v", err)
		}
	}()

	_, err := NewKafka(Config{Broker: "example.com"})
	if err == nil {
		t.Error("expected error")
	}
}

func TestKafka_ConsumerLoop(t *testing.T) {
	b := new(bytes.Buffer)
	err := types.WeatherForecast{
		ForecastedAt:  time.Now(),
		ForecastedFor: time.Now(),
		Location: types.Location{
			Name:      "Somewhere Hot and Equatorial",
			Latitude:  0.00,
			Longitude: 0.00,
		},
		Temperature: 55,
	}.Marshall(b)

	if err != nil {
		t.Fatal(err)
	}

	data := b.Bytes()

	for _, test := range []struct {
		name        string
		msg         []byte
		expect      string
		expectError bool
	}{
		{"Empty input", []byte(""), "", true},
		{"Valid data", data, "Somewhere Hot and Equatorial", false},
	} {
		t.Run(test.name, func(t *testing.T) {
			k := Kafka{consumer: dummyKafka{message: test.msg, topic: "weather"}}

			c := make(chan Message)
			go func() {
				err := k.ConsumerLoop(c)
				if err == nil && test.expectError {
					t.Errorf("expected error")
				}

				if err != nil && !test.expectError {
					t.Errorf("unexpected error: %+v", err)
				}
			}()

			if !test.expectError {
				i := <-c

				rcvd := i.ct.(*types.WeatherForecast).Location.Name
				if test.expect != rcvd {
					t.Errorf("expected %q, received %q", test.expect, rcvd)
				}
			}
		})
	}
}
