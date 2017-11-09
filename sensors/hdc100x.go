package sensors

type HDC100xSensor struct {
	device iioDevice
}

func NewHDC100xSensor(device iioDevice) HDC100xSensor {
	return HDC100xSensor{device: device}
}

func (s *HDC100xSensor) TemperatureHumidity() (temperature, humidity float64, err error) {
	temperature, err = s.Temperature()
	if err != nil {
		return
	}
	humidity, err = s.Humidity()
	if err != nil {
		return
	}
	return
}

func (s *HDC100xSensor) temperatureDatasheet() (temperature float64, err error) {
	tempRaw, err := s.device.readFloat("in_temp_raw")
	if err != nil {
		return
	}
	temperature = (tempRaw/65536.0)*165.0 - 40.0
	return
}

func (s *HDC100xSensor) temperatureComputed() (temperature float64, err error) {
	tempRaw, err := s.device.readFloat("in_temp_raw")
	if err != nil {
		return
	}
	offset, err := s.device.readFloat("in_temp_offset")
	if err != nil {
		return
	}
	scale, err := s.device.readFloat("in_temp_scale")
	if err != nil {
		return
	}

	temperature = ((tempRaw + offset) * scale) / 1000.0
	return
}

func (s *HDC100xSensor) Temperature() (temperature float64, err error) {
	temperature, err = s.temperatureComputed()
	if err != nil {
		return
	}
	return
}

func (s *HDC100xSensor) humidityDatasheet() (humidity float64, err error) {
	humiRaw, err := s.device.readFloat("in_humidityrelative_raw")
	if err != nil {
		return
	}
	humidity = (humiRaw / 65536.0) * 100.0
	return
}

func (s *HDC100xSensor) humidityComputed() (humidity float64, err error) {
	humiRaw, err := s.device.readFloat("in_humidityrelative_raw")
	if err != nil {
		return
	}
	scale, err := s.device.readFloat("in_humidityrelative_scale")
	if err != nil {
		return
	}
	humidity = humiRaw * scale
	return
}

func (s *HDC100xSensor) Humidity() (humidity float64, err error) {
	humidity, err = s.humidityComputed()
	if err != nil {
		return
	}
	return
}
