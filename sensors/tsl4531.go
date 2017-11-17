package sensors

import "github.com/sensebox/senseboxpi/hardware/iio"
import "github.com/sensebox/senseboxpi/hardware"

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

// TSL4531Lux reads and returns the current light intensity in lux
func (t TSL4531) Lux() (lux float64, err error) {
	lux, err = t.device.ReadFloat("in_illuminance_raw")
	if err != nil {
		return
	}
	return
}

// Pheonomenons returns "bmp280_pressure" for this sensor
func (t TSL4531) Phenomenons() []string {
	return []string{"light"}
}

// ReadValue reads and returns the current atmospheric pressure in hPa
func (t TSL4531) ReadValue(phenomenon string) (float64, error) {
	return t.Lux()
}
