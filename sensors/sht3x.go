package sensors

import (
	"fmt"

	"github.com/sensebox/senseboxpi/hardware"
	"github.com/sensebox/senseboxpi/hardware/hwmon"
)

// SHT3x wraps the Industrial I/O sensor sht3x. Selected by sensorType
// "sht3x" and phenomenons "temperature" and "humidity"
type SHT3x struct {
	device hardware.DeviceI
}

// NewSHT3xSensor initializes a new SensorDeviceI of type SHT3x
func NewSHT3xSensor() (SensorI, error) {
	device, err := hwmon.DeviceByName("sht3x")
	if err != nil {
		return SHT3x{}, err
	}

	return SHT3x{device}, nil
}

// Temperature reads and returns the current temperature in degrees celsius
func (s SHT3x) Temperature() (temperature float64, err error) {
	tempRaw, err := s.device.ReadFloat("temp1_input")
	if err != nil {
		return 0, err
	}
	return tempRaw / 1000.0, nil
}

// Humidity reads and returns the current relative humidity in percent
func (s SHT3x) Humidity() (humidity float64, err error) {
	humRaw, err := s.device.ReadFloat("humidity1_input")
	if err != nil {
		return 0, err
	}
	return humRaw / 1000.0, nil
}

// Phenomenons returns []string{"pressure"} for this sensor
func (s SHT3x) Phenomenons() []string {
	return []string{"temperature", "humidity"}
}

// ReadValue reads and returns the current atmospheric pressure in hPa for
// phenomenon "pressure"
func (s SHT3x) ReadValue(phenomenon string) (float64, error) {
	switch phenomenon {
	case "temperature":
		return s.Temperature()
	case "humidity":
		return s.Humidity()
	}

	return 0, fmt.Errorf("invalid phenomenon %s", phenomenon)
}
