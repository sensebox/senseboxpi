package sensebox

import (
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/parnurzeal/gorequest"
)

const (
	baseURL = "https://api.osem.vo1d.space/boxes/"
)

type SenseBoxSensor struct {
	ID           string `json:"_id"`
	measurements []measurement
}

type SenseBox struct {
	ID      string            `json:"_id"`
	Sensors []*SenseBoxSensor `json:"sensors"`
}

type measurement struct {
	Sensor    *SenseBoxSensor `json:"sensor"`
	Value     Number          `json:"value"`
	Timestamp time.Time       `json:"createdAt,omitempty"`
}

func (s SenseBoxSensor) MarshalJSON() ([]byte, error) {
	return []byte("\"" + s.ID + "\""), nil
}

type Number float64

func (f Number) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", f)), nil
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

func NewSenseBox(ID string, sensors ...*SenseBoxSensor) (SenseBox, error) {
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

	return SenseBox{ID, sensors}, nil
}

func (s *SenseBoxSensor) AddMeasurement(value float64, timestamp time.Time) {
	s.measurements = append(s.measurements, measurement{s, Number(value), timestamp.UTC()})
}

func (s *SenseBox) SubmitMeasurements() []error {
	var measurements []measurement
	for _, sensor := range s.Sensors {
		measurements = append(measurements, sensor.measurements...)
	}
	resp, body, errs := gorequest.New().Post(baseURL + s.ID + "/data").
		Send(measurements).
		End()

	if errs != nil {
		return errs
	}

	fmt.Println(resp.Status, body)

	return nil
}
