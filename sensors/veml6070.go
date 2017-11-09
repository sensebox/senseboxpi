package sensors

type VEML6070Sensor struct {
	device iioDevice
}

func NewVEML6070Sensor(device iioDevice) VEML6070Sensor {
	return VEML6070Sensor{device: device}
}

func (s *VEML6070Sensor) UV() (uv float64, err error) {
	uv, err = s.device.readFloat("in_intensity_uv_raw")
	if err != nil {
		return
	}
	return
}
