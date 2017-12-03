package sensebox

import (
	"fmt"
	"strings"
)

// MarshalJSON of a Sensor returns the ID of the sensor
func (s sensor) MarshalJSON() ([]byte, error) {
	return []byte("\"" + s.ID + "\""), nil
}

type number float64

// MarshalJSON of type number formats the float with two decimals and trims
// excess zeroes and dots
func (f number) MarshalJSON() ([]byte, error) {
	return []byte(f.String()), nil
}

// String returns the numbers string representation (float with two decimals and
// trimed excess zeroes and dot
func (f number) String() string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", f), "0"), ".")
}

// UnmarshalJSON of type id checks the id for validity
func (i *id) UnmarshalJSON(jsonBytes []byte) error {
	rawID := string(jsonBytes[1 : len(jsonBytes)-1])
	if err := validateID(rawID); err != nil {
		return err
	}

	*i = id(rawID)

	return nil
}

// String returns the id as 24 digit hex string
func (i *id) String() string {
	return string(*i)
}
