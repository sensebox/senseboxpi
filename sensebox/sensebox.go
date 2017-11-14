package sensebox

import (
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/parnurzeal/gorequest"
	"github.com/sensebox/senseboxpi/sensors"
)

const (
	baseURL = "https://api.osem.vo1d.space/boxes/"
)

type SenseBoxSensor struct {
	ID     string `json:"_id"`
	Sensor sensors.SensorDevice
	box    *SenseBox
}

type SenseBox struct {
	ID           string           `json:"_id"`
	Sensors      []SenseBoxSensor `json:"sensors"`
	measurements []measurement
}

type measurement struct {
	Sensor    *SenseBoxSensor `json:"sensor"`
	Value     float64         `json:"value"`
	Timestamp time.Time       `json:"createdAt,omitempty"`
}

func (s SenseBoxSensor) MarshalJSON() ([]byte, error) {
	return []byte("\"" + s.ID + "\""), nil
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

	box := SenseBox{ID: ID}

	for _, sensor := range sensors {
		err := validateID(sensor.ID)
		sensor.box = &box
		if err != nil {
			return SenseBox{}, errors.New("Sensor ID " + sensor.ID + " is invalid: " + err.Error())
		}
	}

	box.Sensors = sensors

	return box, nil
}

func (s *SenseBoxSensor) AddMeasurement(value float64, timestamp time.Time) {
	s.box.measurements = append(s.box.measurements, measurement{s, value, timestamp.UTC()})
}

func (s *SenseBox) SubmitMeasurements() []error {
	resp, body, errs := gorequest.New().Post(baseURL + s.ID + "/data").
		Send(s.measurements).
		End()

	if errs != nil {
		return errs
	}

	fmt.Println(resp.Status, body)

	return nil
}
