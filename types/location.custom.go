package types

import (
	"fmt"
)

func (sf Location) ValidLat(s string, v float32) error {
	if sf.Latitude < -90 || sf.Latitude > 90 {
		return fmt.Errorf("%s should be between -90 and 90")
	}

	return nil
}

func (sf Location) ValidLong(s string, v float32) error {
	if sf.Longitude < -180 || sf.Longitude > 180 {
		return fmt.Errorf("%s should be between -90 and 90")
	}

	return nil
}
