package iio

import (
	"github.com/sensebox/senseboxpi/hardware"
)

const iioDir = "/sys/bus/iio/devices/"

// DeviceByName finds the device with the specified name in /sys/bus/iio/devices/
func DeviceByName(name string) (hardware.DeviceI, error) {
	return hardware.SysFsDeviceByName(name, iioDir)
}
