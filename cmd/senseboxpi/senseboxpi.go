package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sensebox/senseboxpi/sensebox"
	"github.com/sensebox/senseboxpi/sensors"
)

var (
	hdc1008, tsl45315, veml6070, bmp280         sensors.SensorDevice
	hdcSensor, tslSensor, vemlSensor, bmpSensor sensebox.SenseBoxSensor
)

func initSensors() (err error) {
	hdc1008, err = sensors.NewHDC100xSensor()
	if err != nil {
		return
	}
	tsl45315, err = sensors.NewTSL4531Sensor()
	if err != nil {
		return
	}
	veml6070, err = sensors.NewVEML6070Sensor()
	if err != nil {
		return
	}
	bmp280, err = sensors.NewBMP280Sensor()
	if err != nil {
		return
	}
	return nil
}

func readSensors() (temperature, humidity, lux, uv, pressure float64, err error) {
	temperature, humidity, err = hdc1008.HDC100xTemperatureHumidity()
	if err != nil {
		return
	}
	lux, err = tsl45315.TSL4531Lux()
	if err != nil {
		return
	}
	uv, err = veml6070.VEML6070UV()
	if err != nil {
		return
	}
	pressure, err = bmp280.BMP280Pressure()
	if err != nil {
		return
	}
	return
}

func initSenseBox() (sensebox.SenseBox, error) {
	hdcSensor = sensebox.SenseBoxSensor{ID: "", Sensor: hdc1008}
	tslSensor = sensebox.SenseBoxSensor{ID: "", Sensor: tsl45315}
	vemlSensor = sensebox.SenseBoxSensor{ID: "", Sensor: veml6070}
	bmpSensor = sensebox.SenseBoxSensor{ID: "", Sensor: bmp280}

	box, err := sensebox.NewSenseBox("", hdcSensor, tslSensor, vemlSensor, bmpSensor)
	if err != nil {
		return sensebox.SenseBox{}, err
	}
	return box, nil
}

func main() {
	err := initSensors()
	if err != nil {
		log.Fatal(err)
	}

	senseBox, err := initSenseBox()
	if err != nil {
		log.Fatal(err)
	}

	temperature, humidity, lux, uv, pressure, err := readSensors()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(temperature)
	fmt.Println(humidity)
	fmt.Println(pressure)
	fmt.Println(lux)
	fmt.Println(uv)

	senseBox.AddMeasurement(&hdcSensor, temperature, time.Now())

	senseBox.SubmitMeasurements()

}
