package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/sensebox/senseboxpi/sensebox"
)

const versionStr = "1.1.0"

func readFlags() (configPath string) {
	const (
		defaultConfigPath = "senseboxpi_config.json"
		usage             = "path of the configuration json"
	)
	flag.StringVar(&configPath, "config", defaultConfigPath, usage)
	flag.StringVar(&configPath, "c", defaultConfigPath, usage+" (shorthand)")
	flag.Parse()

	if flag.Arg(0) == "version" {
		fmt.Printf("senseboxpi version %s\nUsage:\n", versionStr)
		flag.PrintDefaults()
		os.Exit(0)
	}
	return configPath
}

func main() {
	configPath := readFlags()

	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	senseBox, err := sensebox.NewFromJSON(configBytes)
	if err != nil {
		log.Fatal(err)
	}
	errs := senseBox.ReadSensorsAndSubmitMeasurements()
	if errs != nil {
		for _, err := range errs {
			fmt.Println(err)
		}
	}
}
