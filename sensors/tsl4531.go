package sensors

import "github.com/sensebox/senseboxpi/sensors/iio"

// NewTSL4531Sensor initializes a new SensorDevice of type TSL4531
func NewTSL4531Sensor() (SensorDevice, error) {
	device, err := iio.DeviceByName("tsl4531")
	if err != nil {
		return SensorDevice{}, err
	}

	return SensorDevice{device}, nil
}

// TSL4531Lux reads and returns the current light intensity in lux
func (s *SensorDevice) TSL4531Lux() (lux float64, err error) {
	lux, err = s.device.ReadFloat("in_illuminance_raw")
	if err != nil {
		return
	}
	return
}
