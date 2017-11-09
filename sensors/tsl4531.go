package sensors

type TSL4531Sensor struct {
	device iioDevice
}

func NewTSL4531Sensor(device iioDevice) TSL4531Sensor {
	return TSL4531Sensor{device: device}
}

func (s *TSL4531Sensor) Lux() (lux float64, err error) {
	lux, err = s.device.readFloat("in_illuminance_raw")
	if err != nil {
		return
	}
	return
}
