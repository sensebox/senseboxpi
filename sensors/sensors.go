package sensors

import "fmt"

// SensorI defines the methods Sensors should implement
type SensorI interface {
	// Phenomenons should return a string slice of availiable phenomenons of the
	// sensor
	Phenomenons() []string
	// ReadValue should return a sensor reading of the requested phenomenon
	ReadValue(string) (float64, error)
}

type fn func() (SensorI, error)

var initMap = map[string]fn{
	"bmp280":   NewBMP280Sensor,
	"hdc100x":  NewHDC100xSensor,
	"tsl4531":  NewTSL4531Sensor,
	"veml6070": NewVEML6070Sensor,
}

// NewSensor initializes a new SensorDevice from a sensorType and a phenomenon
func NewSensor(sensorType, phenomenon string) (SensorI, error) {
	if initFunc, ok := initMap[sensorType]; ok {
		sensor, err := initFunc()
		if err != nil {
			return nil, err
		}

		for _, sensorPhenomenon := range sensor.Phenomenons() {
			if sensorPhenomenon == phenomenon {
				return sensor, nil
			}
		}
	}

	return nil, fmt.Errorf("no hardware for sensorType \"%s\" and phenomenon \"%s\" found", sensorType, phenomenon)
}
