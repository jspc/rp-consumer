package types

import (
	mint "github.com/vinyl-linux/mint"
	"io"
)

type Location struct {
	// Name contains the name of a location
	Name string
	// Latitude relates the latitude of the described location
	Latitude float32
	// Longitude relates the longitude of the described location
	Longitude float32
}

func (sf Location) Validate() error {
	errors := make([]error, 0)
	for _, err := range []error{mint.StringNotEmpty("Name", sf.Name), sf.ValidLat("Latitude", sf.Latitude), sf.ValidLong("Longitude", sf.Longitude)} {
		if err != nil {
			errors = append(errors, err)
		}
	}
	return mint.ValidationErrors("Location", errors)
}
func (sf *Location) Transform() (err error) {
	return
}
func (sf Location) Value() any {
	return sf
}
func (sf *Location) unmarshallName(r io.Reader) (err error) {
	f := mint.NewStringScalar("")
	err = f.Unmarshall(r)
	if err != nil {
		return
	}
	sf.Name = f.Value().(string)
	return
}
func (sf *Location) unmarshallLatitude(r io.Reader) (err error) {
	f := mint.NewFloat32Scalar(float32(0))
	err = f.Unmarshall(r)
	if err != nil {
		return
	}
	sf.Latitude = f.Value().(float32)
	return
}
func (sf *Location) unmarshallLongitude(r io.Reader) (err error) {
	f := mint.NewFloat32Scalar(float32(0))
	err = f.Unmarshall(r)
	if err != nil {
		return
	}
	sf.Longitude = f.Value().(float32)
	return
}
func (sf *Location) Unmarshall(r io.Reader) (err error) {
	if err = sf.unmarshallName(r); err != nil {
		return
	}
	if err = sf.unmarshallLatitude(r); err != nil {
		return
	}
	if err = sf.unmarshallLongitude(r); err != nil {
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
func (sf Location) Marshall(w io.Writer) (err error) {
	if err = sf.Transform(); err != nil {
		return
	}
	if err = sf.Validate(); err != nil {
		return
	}
	if err = mint.NewStringScalar(sf.Name).Marshall(w); err != nil {
		return
	}
	if err = mint.NewFloat32Scalar(sf.Latitude).Marshall(w); err != nil {
		return
	}
	if err = mint.NewFloat32Scalar(sf.Longitude).Marshall(w); err != nil {
		return
	}
	return
}
