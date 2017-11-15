package sensors

import "github.com/sensebox/senseboxpi/sensors/iio"

// NewVEML6070Sensor initializes a new SensorDevice of type VEML6070
func NewVEML6070Sensor() (SensorDevice, error) {
	device, err := iio.DeviceByName("veml6070")
	if err != nil {
		return SensorDevice{}, err
	}

	return SensorDevice{device}, nil
}

// VEML6070UV reads and returns the current uv intensity in microWatts per square centimeter
func (s *SensorDevice) VEML6070UV() (uv float64, err error) {
	uv, err = s.device.ReadFloat("in_intensity_uv_raw")
	if err != nil {
		return
	}
	return
}
