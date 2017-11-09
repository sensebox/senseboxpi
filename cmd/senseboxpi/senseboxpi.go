package main

import (
	"fmt"
	"log"

	"github.com/sensebox/senseboxpi/sensors"
)

var (
	hdc1008  sensors.HDC100xSensor
	tsl4531  sensors.TSL4531Sensor
	veml6070 sensors.VEML6070Sensor
	bmp280   sensors.BMP280Sensor
)

func initSensors() {
	devices, err := sensors.ReadDevices()
	if err != nil {
		log.Fatal(err)
	}

	for _, device := range devices {
		switch device.Name {
		case "1-0043":
			hdc1008 = sensors.NewHDC100xSensor(device)
		case "tsl4531":
			tsl4531 = sensors.NewTSL4531Sensor(device)
		case "veml6070":
			veml6070 = sensors.NewVEML6070Sensor(device)
		case "bmp280":
			bmp280 = sensors.NewBMP280Sensor(device)
		}
	}
}

func readSensors() (temperature, humidity, lux, uv, pressure float64, err error) {
	temperature, humidity, err = hdc1008.TemperatureHumidity()
	if err != nil {
		return
	}
	lux, err = tsl4531.Lux()
	if err != nil {
		return
	}
	uv, err = veml6070.UV()
	if err != nil {
		return
	}
	pressure, err = bmp280.Pressure()
	if err != nil {
		return
	}
	return
}

func main() {
	initSensors()

	temperature, humidity, lux, uv, pressure, err := readSensors()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(temperature)
	fmt.Println(humidity)
	fmt.Println(pressure)
	fmt.Println(lux)
	fmt.Println(uv)

}
