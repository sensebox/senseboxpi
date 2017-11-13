package iio

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

const iioDir = "/sys/bus/iio/devices/"

type Device struct {
	path string
}

func (d Device) Read(filename string) (result string, err error) {
	resultBytes, err := ioutil.ReadFile(d.path + "/" + filename)
	if err != nil {
		return
	}
	result = strings.TrimSpace(string(resultBytes))
	return
}

func (d Device) ReadFloat(filename string) (result float64, err error) {
	resultStr, err := d.Read(filename)
	if err != nil {
		return
	}
	result, err = strconv.ParseFloat(resultStr, 64)
	return
}

func (d Device) Name() (name string, err error) {
	name, err = d.Read("name")
	if err != nil {
		return
	}
	return
}

var iioDevices []Device

// Devices enumerates Industrial I/O devices located in /sys/bus/iio/devices/
func Devices() ([]Device, error) {
	if len(iioDevices) == 0 {
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
	return iioDevices, nil
}

// DeviceByName finds the device with the specified name in /sys/bus/iio/devices/
func DeviceByName(name string) (Device, error) {
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
		//switch device.Name {
		//case "1-0043":
		//	hdc1008 = sensors.NewHDC100xSensor(device)
		//case "tsl4531":
		//	tsl45315 = sensors.NewTSL4531Sensor(device)
		//case "veml6070":
		//	veml6070 = sensors.NewVEML6070Sensor(device)
		//case "bmp280":
		//	bmp280 = sensors.NewBMP280Sensor(device)
		//}
	}

	return Device{}, errors.New("no device with name " + name + "availiable")

}
