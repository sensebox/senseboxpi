package sensors

type BMP280Sensor struct {
	device iioDevice
}

func NewBMP280Sensor(device iioDevice) BMP280Sensor {
	return BMP280Sensor{device: device}
}

func (s *BMP280Sensor) Pressure() (pressure float64, err error) {
	pressureRaw, err := s.device.readFloat("in_pressure_input")
	if err != nil {
		return
	}
	pressure = pressureRaw * 10.0
	return
}
