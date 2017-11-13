package sensebox

import (
	"encoding/hex"
	"errors"
	"time"

	"github.com/sensebox/senseboxpi/sensors"
)

const (
	baseURL = "https://api.osem.vo1d.space/boxes"
)

type SenseBoxSensor struct {
	ID     string `json:"_id"`
	Sensor sensors.SensorDevice
}

type SenseBox struct {
	ID           string           `json:"_id"`
	Sensors      []SenseBoxSensor `json:"sensors"`
	measurements []measurement
}

type measurement struct {
	sensor    string    `json:"sensor"`
	value     float64   `json:"value"`
	timestamp time.Time `json:"createdAt":omitempty`
}

func validateID(id string) (err error) {
	hex, err := hex.DecodeString(id)
	if err != nil {
		return
	}

	if len(hex) != 12 {
		return errors.New("id must be exactly 24 characters long")
	}

	return
}

func NewSenseBox(ID string, sensors ...SenseBoxSensor) (SenseBox, error) {
	err := validateID(ID)
	if err != nil {
		return SenseBox{}, errors.New("senseBoxID " + ID + " is invalid: " + err.Error())
	}

	for _, sensor := range sensors {
		err := validateID(sensor.ID)
		if err != nil {
			return SenseBox{}, errors.New("Sensor ID " + sensor.ID + " is invalid: " + err.Error())
		}
	}

	return SenseBox{ID: ID, Sensors: sensors}, nil
}

func (s *SenseBox) AddMeasurement(sensor *SenseBoxSensor, value float64, timestamp time.Time) {
	s.measurements = append(s.measurements, measurement{sensor.ID, value, timestamp})
}

func (s *SenseBox) SubmitMeasurements() {

}
