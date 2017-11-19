package hardware

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// SysFsDevice is used as embedding for iio and hwmon devices
type SysFsDevice struct {
	Path string
}

// Read reads a value from the specified filename.
func (s SysFsDevice) Read(filename string) (result string, err error) {
	resultBytes, err := ioutil.ReadFile(s.Path + "/" + filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(resultBytes)), nil
}

// ReadFloat reads a float64 from the specified filename
func (s SysFsDevice) ReadFloat(filename string) (result float64, err error) {
	resultStr, err := s.Read(filename)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(resultStr, 64)
}

// Name reads the name of the Device from the "name" file
func (s SysFsDevice) name() (name string, err error) {
	return s.Read("name")
}

// EnumerateDevices enumerates devices located in given directory
func enumerateSysFsDevices(path string) ([]SysFsDevice, error) {
	var devices []SysFsDevice
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		d := SysFsDevice{path + file.Name()}
		devices = append(devices, d)
	}
	return devices, nil
}

// SysFsDeviceByName looks for the device with the given name in the given path
func SysFsDeviceByName(name, path string) (DeviceI, error) {
	devices, err := enumerateSysFsDevices(path)
	if err != nil {
		return SysFsDevice{}, err
	}

	for _, device := range devices {
		// ignore error, because we don't want to exit the loop prematurely
		deviceName, _ := device.name()
		if deviceName == name {
			return device, nil
		}
	}

	return SysFsDevice{}, fmt.Errorf("no device with name %s in path %s available", name, path)
}
