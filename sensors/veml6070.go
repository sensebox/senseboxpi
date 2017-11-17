package sensors

import "github.com/sensebox/senseboxpi/hardware/iio"
import "github.com/sensebox/senseboxpi/hardware"

type VEML6070 struct {
	device hardware.DeviceI
}

// NewVEML6070Sensor initializes a new SensorDevice of type VEML6070
func NewVEML6070Sensor() (SensorI, error) {
	device, err := iio.DeviceByName("veml6070")
	if err != nil {
		return VEML6070{}, err
	}

	return VEML6070{device}, nil
}

// VEML6070UV reads and returns the current uv intensity in microWatts per square centimeter
func (v VEML6070) UV() (uv float64, err error) {
	uv, err = v.device.ReadFloat("in_intensity_uv_raw")
	if err != nil {
		return
	}
	return
}

func (v VEML6070) Phenomenons() []string {
	return []string{"uv"}
}

// ReadValue reads and returns the current atmospheric pressure in hPa
func (v VEML6070) ReadValue(phenomenon string) (float64, error) {
	return v.UV()
}
