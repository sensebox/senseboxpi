package sensors

import (
	"github.com/sensebox/senseboxpi/sensors/iio"
)

func NewBMP280Sensor() (SensorDevice, error) {
	device, err := iio.DeviceByName("bmp280")
	if err != nil {
		return SensorDevice{}, err
	}

	return SensorDevice{device}, nil
}

func (s *SensorDevice) BMP280Pressure() (pressure float64, err error) {
	pressureRaw, err := s.device.ReadFloat("in_pressure_input")
	if err != nil {
		return
	}
	pressure = pressureRaw * 10.0
	return
}
