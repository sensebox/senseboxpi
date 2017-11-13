package sensors

import "github.com/sensebox/senseboxpi/sensors/iio"

func NewTSL4531Sensor() (SensorDevice, error) {
	device, err := iio.DeviceByName("tsl4531")
	if err != nil {
		return SensorDevice{}, err
	}

	return SensorDevice{device}, nil
}

func (s *SensorDevice) TSL4531Lux() (lux float64, err error) {
	lux, err = s.device.ReadFloat("in_illuminance_raw")
	if err != nil {
		return
	}
	return
}
