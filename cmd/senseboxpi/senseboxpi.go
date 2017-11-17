package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/sensebox/senseboxpi/sensebox"
)

func main() {
	configBytes, err := ioutil.ReadFile("config.json")

	senseBox, err := sensebox.NewFromJSON(configBytes)
	if err != nil {
		log.Fatal(err)
	}
	errs := senseBox.ReadSensorsAndSubmitMeasurements()
	if errs != nil {
		for _, err := range errs {
			fmt.Println(err)
		}
		log.Fatal("")
	}
}
