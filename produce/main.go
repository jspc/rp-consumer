package main

import (
	"math/rand"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/jspc/rp-consumer/types"
)

var loc = types.Location{
	Name:      "Gringley on the Hill",
	Latitude:  53.408395,
	Longitude: -0.8941056,
}

func main() {
	k, err := NewKafka("127.0.0.1:40821")
	if err != nil {
		panic(err)
	}

	go k.logs()

	for _, f := range []func() error{
		k.addSensors,
		k.addPrecipitation,
		k.addWeather,
	} {
		err = f()
		if err != nil {
			panic(err)
		}
	}

	// Ensure flush
	time.Sleep(time.Second * 10)
}

func (k Kafka) addSensors() (err error) {
	t := time.Date(2023, time.September, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	tpc := "atmosphere"

	sensor := uuid.Must(uuid.NewV4())

	for {
		if t.After(now) {
			return
		}

		err = k.Write(tpc, &types.SensorReading{
			Timestamp:   t,
			Location:    loc,
			Temperature: rand.Float32()*10 + 25,
			Humidity:    rand.Float32() * 100,
			PM2_5:       rand.Float32() * 500,
			Sensor:      sensor,
		})
		if err != nil {
			return
		}

		t = t.Add(time.Minute * 15)
	}
}

func (k Kafka) addPrecipitation() (err error) {
	t := time.Date(2023, time.September, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	tpc := "precipitation"

	sensor := uuid.Must(uuid.NewV4())
	precipitation := 0

	for {
		if t.After(now) {
			return
		}

		precipitation += rand.Intn(10)
		err = k.Write(tpc, &types.Precipitation{
			Timestamp:     t,
			Location:      loc,
			Sensor:        sensor,
			Precipitation: float32(precipitation),
		})
		if err != nil {
			return
		}

		t = t.Add(time.Hour)
	}
}

func (k Kafka) addWeather() (err error) {
	t := time.Date(2023, time.September, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	tpc := "weather"

	for {
		if t.After(now) {
			return
		}

		err = k.Write(tpc, &types.WeatherForecast{
			ForecastedAt:  time.Now(),
			ForecastedFor: t,
			Location:      loc,
			Temperature:   rand.Float32()*10 + 25,
		})
		if err != nil {
			return
		}

		t = t.Add(time.Hour * 4)
	}
}
