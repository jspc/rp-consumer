package types

import (
	mint "github.com/vinyl-linux/mint"
	"io"
	"time"
)

type WeatherForecast struct {
	// ForecastedAt contains the datetime at which this forecast was created
	ForecastedAt time.Time
	// Location contains a reference to the specified location of this forecast
	Location Location
	// Temperature is a float containing the forecasted temperature temperature
	Temperature float32
	// Date this forecast is for
	ForecastedFor time.Time
}

func (sf WeatherForecast) Validate() error {
	errors := make([]error, 0)
	for _, err := range []error{mint.DateInPast("ForecastedAt", sf.ForecastedAt)} {
		if err != nil {
			errors = append(errors, err)
		}
	}
	return mint.ValidationErrors("WeatherForecast", errors)
}
func (sf *WeatherForecast) Transform() (err error) {
	sf.ForecastedAt, err = mint.DateInUtc(sf.ForecastedAt)
	if err != nil {
		return
	}
	sf.ForecastedFor, err = mint.DateInUtc(sf.ForecastedFor)
	if err != nil {
		return
	}
	return
}
func (sf WeatherForecast) Value() any {
	return sf
}
func (sf *WeatherForecast) unmarshallForecastedAt(r io.Reader) (err error) {
	f := mint.NewDatetimeScalar(time.Time{})
	err = f.Unmarshall(r)
	if err != nil {
		return
	}
	sf.ForecastedAt = f.Value().(time.Time)
	return
}
func (sf *WeatherForecast) unmarshallLocation(r io.Reader) (err error) {
	f := new(Location)
	err = f.Unmarshall(r)
	if err != nil {
		return
	}
	sf.Location = f.Value().(Location)
	return
}
func (sf *WeatherForecast) unmarshallTemperature(r io.Reader) (err error) {
	f := mint.NewFloat32Scalar(float32(0))
	err = f.Unmarshall(r)
	if err != nil {
		return
	}
	sf.Temperature = f.Value().(float32)
	return
}
func (sf *WeatherForecast) unmarshallForecastedFor(r io.Reader) (err error) {
	f := mint.NewDatetimeScalar(time.Time{})
	err = f.Unmarshall(r)
	if err != nil {
		return
	}
	sf.ForecastedFor = f.Value().(time.Time)
	return
}
func (sf *WeatherForecast) Unmarshall(r io.Reader) (err error) {
	if err = sf.unmarshallForecastedAt(r); err != nil {
		return
	}
	if err = sf.unmarshallLocation(r); err != nil {
		return
	}
	if err = sf.unmarshallTemperature(r); err != nil {
		return
	}
	if err = sf.unmarshallForecastedFor(r); err != nil {
		return
	}
	if err = sf.Transform(); err != nil {
		return
	}
	if err = sf.Validate(); err != nil {
		return
	}
	return
}
func (sf WeatherForecast) Marshall(w io.Writer) (err error) {
	if err = sf.Transform(); err != nil {
		return
	}
	if err = sf.Validate(); err != nil {
		return
	}
	if err = mint.NewDatetimeScalar(sf.ForecastedAt).Marshall(w); err != nil {
		return
	}
	if err = sf.Location.Marshall(w); err != nil {
		return
	}
	if err = mint.NewFloat32Scalar(sf.Temperature).Marshall(w); err != nil {
		return
	}
	if err = mint.NewDatetimeScalar(sf.ForecastedFor).Marshall(w); err != nil {
		return
	}
	return
}
