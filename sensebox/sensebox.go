package sensebox

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
)

const (
	baseURL = "https://api.osem.vo1d.space/boxes/"
)

// A Sensor of a SenseBox
type Sensor struct {
	ID           string `json:"_id"`
	measurements []measurement
}

// SenseBox has an ID and Sensors
type SenseBox struct {
	ID      string    `json:"_id"`
	Sensors []*Sensor `json:"sensors"`
}

type measurement struct {
	Sensor    *Sensor   `json:"sensor"`
	Value     number    `json:"value"`
	Timestamp time.Time `json:"createdAt,omitempty"`
}

// MarshalJSON of a Sensor returns the ID of the sensor
func (s Sensor) MarshalJSON() ([]byte, error) {
	return []byte("\"" + s.ID + "\""), nil
}

type number float64

func (f number) MarshalJSON() ([]byte, error) {
	return []byte(strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", f), "0"), ".")), nil
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

// NewSenseBox initializes a SenseBox. It takes the ID and multiple Sensors
// as parameters.
// It validates the ID and also the IDs of the sensors
func NewSenseBox(ID string, sensors ...*Sensor) (SenseBox, error) {
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

// AddMeasurement adds a new measurement to the Sensor
func (s *Sensor) AddMeasurement(value float64, timestamp time.Time) {
	s.measurements = append(s.measurements, measurement{s, number(value), timestamp.UTC()})
}

// SubmitMeasurements tries to send the measurements of the Sensors of the SenseBox
// to the openSenseMap
func (s *SenseBox) SubmitMeasurements() []error {
	var measurements []measurement
	for _, sensor := range s.Sensors {
		measurements = append(measurements, sensor.measurements...)
		// clear measurements
		sensor.measurements = nil
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
