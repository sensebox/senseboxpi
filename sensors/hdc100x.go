package sensors

import "github.com/sensebox/senseboxpi/sensors/iio"

func NewHDC100xSensor() (SensorDevice, error) {
	device, err := iio.DeviceByName("1-0043")
	if err != nil {
		return SensorDevice{}, err
	}

	return SensorDevice{device}, nil
}

func (s *SensorDevice) HDC100xTemperatureHumidity() (temperature, humidity float64, err error) {
	temperature, err = s.HDC100xTemperature()
	if err != nil {
		return
	}
	humidity, err = s.HDC100xHumidity()
	if err != nil {
		return
	}
	return
}

func (s *SensorDevice) HDC100xTemperature() (temperature float64, err error) {
	tempRaw, err := s.device.ReadFloat("in_temp_raw")
	if err != nil {
		return
	}
	offset, err := s.device.ReadFloat("in_temp_offset")
	if err != nil {
		return
	}
	scale, err := s.device.ReadFloat("in_temp_scale")
	if err != nil {
		return
	}

	temperature = ((tempRaw + offset) * scale) / 1000.0
	return
}

func (s *SensorDevice) HDC100xHumidity() (humidity float64, err error) {
	humiRaw, err := s.device.ReadFloat("in_humidityrelative_raw")
	if err != nil {
		return
	}
	scale, err := s.device.ReadFloat("in_humidityrelative_scale")
	if err != nil {
		return
	}
	humidity = humiRaw * scale
	return
}
