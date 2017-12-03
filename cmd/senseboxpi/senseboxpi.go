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

func readFlags() (configPath, csvOutputPath string, offline bool) {
	const (
		defaultConfigPath = "senseboxpi_config.json"
		configUsage       = "path of the configuration json"
		offlineUsage      = "operate offline. Do not upload to server"
		csvOutputUsage    = "path to file where measurements in csv format will be appended"
	)
	// config flag
	flag.StringVar(&configPath, "config", defaultConfigPath, configUsage)
	flag.StringVar(&configPath, "c", defaultConfigPath, configUsage+" (shorthand)")
	// offline flag
	flag.BoolVar(&offline, "offline", false, offlineUsage)
	// csv output path flag
	flag.StringVar(&csvOutputPath, "csv-output", "", csvOutputUsage)

	flag.Parse()

	if flag.Arg(0) == "version" {
		fmt.Printf("senseboxpi version %s\nUsage:\n", versionStr)
		flag.PrintDefaults()
		os.Exit(0)
	}
	return configPath, csvOutputPath, offline
}

func main() {
	configPath, _, _ := readFlags()

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
