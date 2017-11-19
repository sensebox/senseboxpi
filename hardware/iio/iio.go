package iio

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/sensebox/senseboxpi/hardware"
)

const iioDir = "/sys/bus/iio/devices/"

// Device is a Industrial I/O device
type Device struct {
	path string
}

// Read reads a value from the specified filenamehardware.
func (d Device) Read(filename string) (result string, err error) {
	resultBytes, err := ioutil.ReadFile(d.path + "/" + filename)
	if err != nil {
		return
	}
	result = strings.TrimSpace(string(resultBytes))
	return
}

// ReadFloat reads a float64 from the specified filename
func (d Device) ReadFloat(filename string) (result float64, err error) {
	resultStr, err := d.Read(filename)
	if err != nil {
		return
	}
	result, err = strconv.ParseFloat(resultStr, 64)
	return
}

// Name reads the name of the Device from the "name" file
func (d Device) Name() (name string, err error) {
	name, err = d.Read("name")
	if err != nil {
		return
	}
	return
}

// Devices enumerates Industrial I/O devices located in /sys/bus/iio/devices/
func Devices() ([]hardware.DeviceI, error) {
	var iioDevices []hardware.DeviceI
	files, err := ioutil.ReadDir(iioDir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		d := Device{iioDir + file.Name()}
		iioDevices = append(iioDevices, d)
	}
	return iioDevices, nil
}

// DeviceByName finds the device with the specified name in /sys/bus/iio/devices/
func DeviceByName(name string) (hardware.DeviceI, error) {
	devices, err := Devices()
	if err != nil {
		return Device{}, err
	}

	for _, device := range devices {
		// ignore error, because we don't want to exit the loop prematurely
		deviceName, _ := device.Name()
		if deviceName == name {
			return device, nil
		}
	}

	return Device{}, fmt.Errorf("no iio device with name %s available", name)
}
