package types

import (
	v5 "github.com/gofrs/uuid/v5"
	mint "github.com/vinyl-linux/mint"
	"io"
	"time"
)

type Precipitation struct {
	// Timestamp the reading was taken at
	Timestamp time.Time
	// Location contains a reference to the specified location of this forecast
	Location Location
	// Specific sensor used for this reading
	Sensor v5.UUID
	// Precipitation holds the returned precipitaton level in mm
	Precipitation float32
}

func (sf Precipitation) Validate() error {
	errors := make([]error, 0)
	for _, err := range []error{mint.DateInPast("Timestamp", sf.Timestamp)} {
		if err != nil {
			errors = append(errors, err)
		}
	}
	return mint.ValidationErrors("Precipitation", errors)
}
func (sf *Precipitation) Transform() (err error) {
	sf.Timestamp, err = mint.DateInUtc(sf.Timestamp)
	if err != nil {
		return
	}
	return
}
func (sf Precipitation) Value() any {
	return sf
}
func (sf *Precipitation) unmarshallTimestamp(r io.Reader) (err error) {
	f := mint.NewDatetimeScalar(time.Time{})
	err = f.Unmarshall(r)
	if err != nil {
		return
	}
	sf.Timestamp = f.Value().(time.Time)
	return
}
func (sf *Precipitation) unmarshallLocation(r io.Reader) (err error) {
	f := new(Location)
	err = f.Unmarshall(r)
	if err != nil {
		return
	}
	sf.Location = f.Value().(Location)
	return
}
func (sf *Precipitation) unmarshallSensor(r io.Reader) (err error) {
	f := mint.NewUuidScalar(v5.UUID{})
	err = f.Unmarshall(r)
	if err != nil {
		return
	}
	sf.Sensor = f.Value().(v5.UUID)
	return
}
func (sf *Precipitation) unmarshallPrecipitation(r io.Reader) (err error) {
	f := mint.NewFloat32Scalar(float32(0))
	err = f.Unmarshall(r)
	if err != nil {
		return
	}
	sf.Precipitation = f.Value().(float32)
	return
}
func (sf *Precipitation) Unmarshall(r io.Reader) (err error) {
	if err = sf.unmarshallTimestamp(r); err != nil {
		return
	}
	if err = sf.unmarshallLocation(r); err != nil {
		return
	}
	if err = sf.unmarshallSensor(r); err != nil {
		return
	}
	if err = sf.unmarshallPrecipitation(r); err != nil {
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
func (sf Precipitation) Marshall(w io.Writer) (err error) {
	if err = sf.Transform(); err != nil {
		return
	}
	if err = sf.Validate(); err != nil {
		return
	}
	if err = mint.NewDatetimeScalar(sf.Timestamp).Marshall(w); err != nil {
		return
	}
	if err = sf.Location.Marshall(w); err != nil {
		return
	}
	if err = mint.NewUuidScalar(sf.Sensor).Marshall(w); err != nil {
		return
	}
	if err = mint.NewFloat32Scalar(sf.Precipitation).Marshall(w); err != nil {
		return
	}
	return
}
