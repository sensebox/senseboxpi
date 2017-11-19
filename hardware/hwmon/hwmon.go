package hwmon

import (
	"github.com/sensebox/senseboxpi/hardware"
)

const hwmonDir = "/sys/class/hwmon/"

// DeviceByName finds the device with the specified name in /sys/class/hwmon/
func DeviceByName(name string) (hardware.DeviceI, error) {
	return hardware.SysFsDeviceByName(name, hwmonDir)
}
