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

func (f number) MarshalJSON() ([]byte, error) {
	return []byte(strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", f), "0"), ".")), nil
}

func (i id) UnmarshalJSON(jsonBytes []byte) error {
	rawID := string(jsonBytes[1 : len(jsonBytes)-1])
	if err := validateID(rawID); err != nil {
		return err
	}

	i = id(rawID)

	return nil
}
