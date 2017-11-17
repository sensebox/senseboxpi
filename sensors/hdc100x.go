package sensors

import (
	"errors"

	"github.com/sensebox/senseboxpi/hardware"
	"github.com/sensebox/senseboxpi/hardware/iio"
)

type HDC100x struct {
	device hardware.DeviceI
}

// NewHDC100xSensor initializes a new SensorDevice of type HDC100x
func NewHDC100xSensor() (SensorI, error) {
	device, err := iio.DeviceByName("1-0043")
	if err != nil {
		return HDC100x{}, err
	}

	return HDC100x{device}, nil
}

// TemperatureHumidity reads and returns the current temperature in degree celsius
// and relative humidity in percent
func (h HDC100x) TemperatureHumidity() (temperature, humidity float64, err error) {
	temperature, err = h.Temperature()
	if err != nil {
		return
	}
	humidity, err = h.Humidity()
	if err != nil {
		return
	}
	return
}

// HDC100xTemperature reads and returns the current temperature in degrees celsius
func (h HDC100x) Temperature() (temperature float64, err error) {
	tempRaw, err := h.device.ReadFloat("in_temp_raw")
	if err != nil {
		return
	}
	offset, err := h.device.ReadFloat("in_temp_offset")
	if err != nil {
		return
	}
	scale, err := h.device.ReadFloat("in_temp_scale")
	if err != nil {
		return
	}

	temperature = ((tempRaw + offset) * scale) / 1000.0
	return
}

// HDC100x reads and returns the current relative humidity in percent
func (h HDC100x) Humidity() (humidity float64, err error) {
	humiRaw, err := h.device.ReadFloat("in_humidityrelative_raw")
	if err != nil {
		return
	}
	scale, err := h.device.ReadFloat("in_humidityrelative_scale")
	if err != nil {
		return
	}
	humidity = humiRaw * scale
	return
}

// Phenomenons returns "temperature" and "humidity" for this sensor
func (h HDC100x) Phenomenons() []string {
	return []string{"temperature", "humidity"}
}

// ReadValue reads and returns the current atmospheric pressure in hPa
func (h HDC100x) ReadValue(phenomenon string) (float64, error) {
	temperature, humidity, err := h.TemperatureHumidity()
	if err != nil {
		return 0, err
	}
	switch phenomenon {
	case "temperature":
		return temperature, nil
	case "humidity":
		return humidity, nil
	}

	return 0, errors.New("invalid phenomenon")
}
