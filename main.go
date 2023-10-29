package main

import (
	"log"
)

func main() {
	c, err := LoadConfig(".config.toml")
	if err != nil {
		panic(err)
	}

	k, err := NewKafka(c)
	if err != nil {
		panic(err)
	}

	messageChan := make(chan Message)
	go func() {
		panic(k.ConsumerLoop(messageChan))
	}()

	for m := range messageChan {
		log.Printf("Writing %#v to table: %q", m.ct, m.topic)
	}
}
