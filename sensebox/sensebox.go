package sensebox

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

const (
	baseURL = "https://api.osem.vo1d.space/boxes/"
)

type id string

type senseBox struct {
	ID      id        `json:"_id"`
	Sensors []*sensor `json:"sensors"`
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

func NewFromJSON(jsonBytes []byte) (senseBox, error) {
	var sbx senseBox
	err := json.Unmarshal(jsonBytes, &sbx)
	if err != nil {
		return senseBox{}, err
	}

	return sbx, nil
}

// SubmitMeasurements tries to send the measurements of the Sensors of the SenseBox
// to the openSenseMap
func (s *senseBox) SubmitMeasurements() []error {
	var measurements []measurement
	for _, sensor := range s.Sensors {
		measurements = append(measurements, sensor.measurements...)
		// clear measurements
		sensor.measurements = nil
	}
	resp, body, errs := gorequest.New().Post(baseURL + string(s.ID) + "/data").
		Send(measurements).
		End()

	if errs != nil {
		return errs
	}

	fmt.Println(resp.Status, body)

	return nil
}

func (s *senseBox) ReadSensorsAndSubmitMeasurements() []error {
	for _, sensor := range s.Sensors {
		if err := sensor.AddMeasurementReading(); err != nil {
			return []error{err}
		}
	}

	return s.SubmitMeasurements()
}
