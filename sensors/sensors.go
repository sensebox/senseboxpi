package sensors

import "errors"

type SensorI interface {
	Phenomenons() []string
	ReadValue(string) (float64, error)
}

type fn func() (SensorI, error)

var initMap = map[string]fn{
	"bmp280":   NewBMP280Sensor,
	"hdc1008":  NewHDC100xSensor,
	"tsl45315": NewTSL4531Sensor,
	"veml6070": NewVEML6070Sensor,
}

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

	return nil, errors.New("no hardware for sensorType \"" + sensorType + "\" and phenomenon \"" + phenomenon + "\" found")
}
