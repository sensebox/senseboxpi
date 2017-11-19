package sensors

import (
	"fmt"

	"github.com/sensebox/senseboxpi/hardware"
	"github.com/sensebox/senseboxpi/hardware/iio"
)

// TSL4531 wraps the Industrial I/O sensor tsl4531. Selected by sensorType
// "tsl4531" and phenomenon "temperature" or "lux"
type TSL4531 struct {
	device hardware.DeviceI
}

// NewTSL4531Sensor initializes a new SensorDevice of type TSL4531
func NewTSL4531Sensor() (SensorI, error) {
	device, err := iio.DeviceByName("tsl4531")
	if err != nil {
		return TSL4531{}, err
	}

	return TSL4531{device}, nil
}

// Lux reads and returns the current light intensity in lux
func (t TSL4531) Lux() (lux float64, err error) {
	lux, err = t.device.ReadFloat("in_illuminance_raw")
	if err != nil {
		return
	}
	return
}

// Phenomenons returns []string{"light"} for this sensor
func (t TSL4531) Phenomenons() []string {
	return []string{"light"}
}

// ReadValue reads and returns the current light intensity in lux for phenomenon
// "light"
func (t TSL4531) ReadValue(phenomenon string) (float64, error) {
	if phenomenon != "light" {
		return 0, fmt.Errorf("invalid phenomenon %s", phenomenon)
	}

	lux, err := t.Lux()
	if err != nil {
		return 0, err
	}

	return lux, nil
}
