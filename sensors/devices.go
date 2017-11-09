package sensors

import (
	"io/ioutil"
	"strconv"
	"strings"
)

const iioDir = "/sys/bus/iio/devices/"

type iioDevice struct {
	Name string
	path string
}

func (d *iioDevice) read(filename string) (result string, err error) {
	resultBytes, err := ioutil.ReadFile(d.path + "/" + filename)
	if err != nil {
		return
	}
	result = strings.TrimSpace(string(resultBytes))
	return
}

func (d *iioDevice) readInt(filename string) (result int64, err error) {
	resultStr, err := d.read(filename)
	if err != nil {
		return
	}
	result, err = strconv.ParseInt(resultStr, 10, 64)
	return
}

func (d *iioDevice) readFloat(filename string) (result float64, err error) {
	resultStr, err := d.read(filename)
	if err != nil {
		return
	}
	result, err = strconv.ParseFloat(resultStr, 64)
	return
}

func (d *iioDevice) readName() error {
	name, err := d.read("name")
	if err != nil {
		return err
	}
	d.Name = name
	return nil
}

func ReadDevices() (devices []iioDevice, err error) {
	files, err := ioutil.ReadDir(iioDir)
	if err != nil {
		return
	}
	for _, file := range files {
		d := iioDevice{path: iioDir + file.Name()}
		d.readName()
		devices = append(devices, d)
	}
	return
}
