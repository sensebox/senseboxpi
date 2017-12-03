package sensebox

import (
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/parnurzeal/gorequest"
)

type id string

type senseBox struct {
	ID           id        `json:"_id"`
	Sensors      []*sensor `json:"sensors"`
	PostDomain   string    `json:"postDomain"`
	measurements []measurement
}

func validateID(id string) error {
	hex, err := hex.DecodeString(id)
	if err != nil {
		return fmt.Errorf("id \"%s\" is invalid: %s", id, err.Error())
	}

	if len(hex) != 12 {
		return fmt.Errorf("id \"%s\" is invalid: id must be exactly 24 characters long", id)
	}

	return nil
}

// NewFromJSON initializes a new senseBox with sensors and postDomain from given
// byte array encoded json obtained from for example ioutil.ReadFile
func NewFromJSON(jsonBytes []byte) (senseBox, error) {
	var sbx senseBox
	err := json.Unmarshal(jsonBytes, &sbx)
	if err != nil {
		return senseBox{}, err
	}

	return sbx, nil
}

// SubmitMeasurements tries to send the measurements of the Sensors of the senseBox
// to the openSenseMap
func (s *senseBox) SubmitMeasurements() []error {
	if len(s.measurements) == 0 {
		return []error{errors.New("No measurements. Did you forgot to call ReadSensors?")}
	}

	postURL, err := url.Parse("https://" + s.PostDomain + "/boxes/" + string(s.ID) + "/data")
	if err != nil {
		return []error{err}
	}
	resp, body, errs := gorequest.New().Post(postURL.String()).
		Send(s.measurements).
		End()

	if errs != nil {
		return errs
	}

	fmt.Println(resp.Status, body)

	return nil
}

// ClearMeasurements clears the measurements previously read through ReadSensors
func (s *senseBox) ClearMeasurements() {
	s.measurements = make([]measurement, len(s.Sensors))
}

// ReadSensors reads measurements from all sensors
func (s *senseBox) ReadSensors() error {
	for _, sensor := range s.Sensors {
		m, err := sensor.ReadMeasurement()
		if err != nil {
			return err
		}
		s.measurements = append(s.measurements, m)
	}
	return nil
}

// ReadSensorsAndSubmitMeasurements takes readings from all sensors and submits
// these measurements through calling SubmitMeasurements
func (s *senseBox) ReadSensorsAndSubmitMeasurements() []error {
	err := s.ReadSensors()
	if err != nil {
		return []error{err}
	}

	return s.SubmitMeasurements()
}

func (s *senseBox) AppendCSV(path string) error {
	if len(s.measurements) == 0 {
		return errors.New("No measurements. Did you forgot to call ReadSensors?")
	}

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	// csv-stringify the measurements
	for _, measurement := range s.measurements {
		m := []string{measurement.Sensor.ID.String(), measurement.Value.String(), measurement.Timestamp.Format(time.RFC3339)}
		writer.Write(m)
	}

	if err := writer.Error(); err != nil {
		return err
	}
	return nil
}
