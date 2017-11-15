package main

import (
	"log"
	"time"

	"github.com/sensebox/senseboxpi/sensebox"
	"github.com/sensebox/senseboxpi/sensors"
)

func initSenseBox() (sensebox.SenseBox, error) {
	tempSensor := sensebox.Sensor{ID: ""}
	humiSensor := sensebox.Sensor{ID: ""}
	lightSensor := sensebox.Sensor{ID: ""}
	uvSensor := sensebox.Sensor{ID: ""}
	presSensor := sensebox.Sensor{ID: ""}

	box, err := sensebox.NewSenseBox("", &tempSensor, &humiSensor, &lightSensor, &uvSensor, &presSensor)
	if err != nil {
		return sensebox.SenseBox{}, err
	}
	return box, nil
}

func main() {
	var (
		tempSensor, humiSensor, lightSensor, uvSensor, presSensor sensebox.Sensor
	)

	hdc1008, tsl45315, veml6070, bmp280, err := initSensors()
	if err != nil {
		log.Fatal(err)
	}

	senseBox, err := initSenseBox()
	if err != nil {
		log.Fatal(err)
	}

	temperature, humidity, err := hdc1008.HDC100xTemperatureHumidity()
	if err != nil {
		log.Fatal(err)
	}
	tempSensor.AddMeasurement(temperature, time.Now())
	humiSensor.AddMeasurement(humidity, time.Now())

	lux, err := tsl45315.TSL4531Lux()
	if err != nil {
		log.Fatal(err)
	}
	lightSensor.AddMeasurement(lux, time.Now())

	uv, err := veml6070.VEML6070UV()
	if err != nil {
		log.Fatal(err)
	}
	uvSensor.AddMeasurement(uv, time.Now())

	pressure, err := bmp280.BMP280Pressure()
	if err != nil {
		log.Fatal(err)
	}
	presSensor.AddMeasurement(pressure, time.Now())

	errs := senseBox.SubmitMeasurements()
	if errs != nil {
		log.Fatal(errs)
	}

}

func initSensors() (hdc1008, tsl45315, veml6070, bmp280 sensors.SensorDevice, err error) {
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
	return
}
