package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/jspc/rp-consumer/types"
	"github.com/pelletier/go-toml/v2"
)

type UnknownTopicError string

func (e UnknownTopicError) Error() string {
	return fmt.Sprintf("Unknown/ unconfigured topic %q", string(e))
}

type Config struct {
	Broker string
	Topic  map[string]Mapping
}

func (c Config) Topics() (t []string) {
	t = make([]string, 0)

	for k := range c.Topic {
		t = append(t, k)
	}

	// Ensure a predictable output
	sort.Strings(t)

	return
}

func (c Config) Type(topic string) (ct consumerType, err error) {
	switch topic {
	case "atmosphere":
		return new(types.SensorReading), nil

	case "precipitation":
		return new(types.Precipitation), nil

	case "weather":
		return new(types.WeatherForecast), nil

	default:
		return nil, UnknownTopicError(topic)
	}
}

type Mapping struct {
	Type string
}

func LoadConfig(fn string) (c Config, err error) {
	data, err := os.ReadFile(fn)
	if err != nil {
		return
	}

	err = toml.Unmarshal(data, &c)

	return
}
