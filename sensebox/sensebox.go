package sensebox

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/parnurzeal/gorequest"
)

type id string

type senseBox struct {
	ID         id        `json:"_id"`
	Sensors    []*sensor `json:"sensors"`
	PostDomain string    `json:"postDomain"`
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
	var measurements []measurement
	for _, sensor := range s.Sensors {
		measurements = append(measurements, sensor.measurements...)
		// clear measurements
		sensor.measurements = nil
	}
	postURL, err := url.Parse("https://" + s.PostDomain + "/boxes/" + string(s.ID) + "/data")
	if err != nil {
		return []error{err}
	}
	resp, body, errs := gorequest.New().Post(postURL.String()).
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
