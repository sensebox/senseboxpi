package sensors

import (
	"github.com/sensebox/senseboxpi/hardware"
	"github.com/sensebox/senseboxpi/hardware/iio"
)

// BMP280 wraps the Industrial I/O sensor bmp280
type BMP280 struct {
	device hardware.DeviceI
}

// NewBMP280Sensor initializes a new SensorDeviceI of type BMP280
func NewBMP280Sensor() (SensorI, error) {
	device, err := iio.DeviceByName("bmp280")
	if err != nil {
		return BMP280{}, err
	}

	return BMP280{device}, nil
}

// Pressure reads and returns the current atmospheric pressure in hPa
func (b BMP280) Pressure() (pressure float64, err error) {
	pressureRaw, err := b.device.ReadFloat("in_pressure_input")
	if err != nil {
		return
	}
	pressure = pressureRaw * 10.0
	return
}

// Pheonomenons returns "bmp280_pressure" for this sensor
func (b BMP280) Phenomenons() []string {
	return []string{"pressure"}
}

// ReadValue reads and returns the current atmospheric pressure in hPa
func (b BMP280) ReadValue(phenomenon string) (float64, error) {
	return b.Pressure()
}
