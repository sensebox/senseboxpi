package sensors

import (
	"fmt"

	"github.com/sensebox/senseboxpi/hardware"
	"github.com/sensebox/senseboxpi/hardware/iio"
)

// VEML6070 wraps the Industrial I/O sensor veml6070. Selected by sensorType
// "veml6070" and phenomenon "uv"
type VEML6070 struct {
	device hardware.DeviceI
}

// NewVEML6070Sensor initializes a new SensorDevice of type VEML6070
func NewVEML6070Sensor() (SensorI, error) {
	device, err := iio.DeviceByName("veml6070")
	if err != nil {
		return VEML6070{}, err
	}

	return VEML6070{device}, nil
}

// UV reads and returns the current uv intensity in microWatts per square
// centimeter
func (v VEML6070) UV() (uv float64, err error) {
	uv, err = v.device.ReadFloat("in_intensity_uv_raw")
	if err != nil {
		return
	}
	return
}

// Phenomenons returns []string{"uv"} for this sensor
func (v VEML6070) Phenomenons() []string {
	return []string{"uv"}
}

// ReadValue reads and returns the current uv intensity in microWatts per square
// centimeter
func (v VEML6070) ReadValue(phenomenon string) (float64, error) {
	if phenomenon != "uv" {
		return 0, fmt.Errorf("invalid phenomenon %s", phenomenon)
	}

	uv, err := v.UV()
	if err != nil {
		return 0, err
	}

	return uv, nil
}
