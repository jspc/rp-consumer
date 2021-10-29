package types

import (
	v5 "github.com/gofrs/uuid/v5"
	mint "github.com/vinyl-linux/mint"
	"io"
	"time"
)

type SensorReading struct {
	// Timestamp the reading was taken at
	Timestamp time.Time
	// Location contains a reference to the specified location of this forecast
	Location Location
	// Temperature contained in this reading
	Temperature float32
	// Humidity contained in this reading
	Humidity float32
	// PM2_5 of the air in this reading
	PM2_5 float32
	// Specific sensor used for this reading
	Sensor v5.UUID
}

func (sf SensorReading) Validate() error {
	errors := make([]error, 0)
	for _, err := range []error{mint.DateInPast("Timestamp", sf.Timestamp)} {
		if err != nil {
			errors = append(errors, err)
		}
	}
	return mint.ValidationErrors("SensorReading", errors)
}
func (sf *SensorReading) Transform() (err error) {
	sf.Timestamp, err = mint.DateInUtc(sf.Timestamp)
	if err != nil {
		return
	}
	return
}
func (sf SensorReading) Value() any {
	return sf
}
func (sf *SensorReading) unmarshallTimestamp(r io.Reader) (err error) {
	f := mint.NewDatetimeScalar(time.Time{})
	err = f.Unmarshall(r)
	if err != nil {
		return
	}
	sf.Timestamp = f.Value().(time.Time)
	return
}
func (sf *SensorReading) unmarshallLocation(r io.Reader) (err error) {
	f := new(Location)
	err = f.Unmarshall(r)
	if err != nil {
		return
	}
	sf.Location = f.Value().(Location)
	return
}
func (sf *SensorReading) unmarshallTemperature(r io.Reader) (err error) {
	f := mint.NewFloat32Scalar(float32(0))
	err = f.Unmarshall(r)
	if err != nil {
		return
	}
	sf.Temperature = f.Value().(float32)
	return
}
func (sf *SensorReading) unmarshallHumidity(r io.Reader) (err error) {
	f := mint.NewFloat32Scalar(float32(0))
	err = f.Unmarshall(r)
	if err != nil {
		return
	}
	sf.Humidity = f.Value().(float32)
	return
}
func (sf *SensorReading) unmarshallPM2_5(r io.Reader) (err error) {
	f := mint.NewFloat32Scalar(float32(0))
	err = f.Unmarshall(r)
	if err != nil {
		return
	}
	sf.PM2_5 = f.Value().(float32)
	return
}
func (sf *SensorReading) unmarshallSensor(r io.Reader) (err error) {
	f := mint.NewUuidScalar(v5.UUID{})
	err = f.Unmarshall(r)
	if err != nil {
		return
	}
	sf.Sensor = f.Value().(v5.UUID)
	return
}
func (sf *SensorReading) Unmarshall(r io.Reader) (err error) {
	if err = sf.unmarshallTimestamp(r); err != nil {
		return
	}
	if err = sf.unmarshallLocation(r); err != nil {
		return
	}
	if err = sf.unmarshallTemperature(r); err != nil {
		return
	}
	if err = sf.unmarshallHumidity(r); err != nil {
		return
	}
	if err = sf.unmarshallPM2_5(r); err != nil {
		return
	}
	if err = sf.unmarshallSensor(r); err != nil {
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
func (sf SensorReading) Marshall(w io.Writer) (err error) {
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
	if err = mint.NewFloat32Scalar(sf.Temperature).Marshall(w); err != nil {
		return
	}
	if err = mint.NewFloat32Scalar(sf.Humidity).Marshall(w); err != nil {
		return
	}
	if err = mint.NewFloat32Scalar(sf.PM2_5).Marshall(w); err != nil {
		return
	}
	if err = mint.NewUuidScalar(sf.Sensor).Marshall(w); err != nil {
		return
	}
	return
}
