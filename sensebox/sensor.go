package sensebox

import (
	"time"

	"github.com/sensebox/senseboxpi/sensors"
)

// A Sensor of a SenseBox
type sensor struct {
	ID           id     `json:"_id"`
	Phenomenon   string `json:"phenomenon"`
	SensorType   string `json:"sensorType"`
	measurements []measurement
	sensorDevice sensors.SensorI
}

type measurement struct {
	Sensor    *sensor   `json:"sensor"`
	Value     number    `json:"value"`
	Timestamp time.Time `json:"createdAt,omitempty"`
}

// InitializeDevice tries to initialize a new SensorDevice from the sensors
// SensorType and Phenomenon
func (s *sensor) InitializeDevice() error {
	if s.sensorDevice == nil {
		sensorDevice, err := sensors.NewSensor(s.SensorType, s.Phenomenon)
		if err != nil {
			return err
		}
		s.sensorDevice = sensorDevice
	}
	return nil
}

// TakeReading tries to take a reading from the sensors sensorDevice
// ReadValue(string). Calls InitializeDevice if the sensorDevice is nil
func (s *sensor) TakeReading() (float64, error) {
	if s.sensorDevice == nil {
		err := s.InitializeDevice()
		if err != nil {
			return 0, err
		}
	}
	return s.sensorDevice.ReadValue(s.Phenomenon)
}

// AddMeasurementReading calls the sensors TakeReading function and adds the
//result to the sensors measurements through AddMeasurement
func (s *sensor) AddMeasurementReading() error {
	reading, err := s.TakeReading()
	if err != nil {
		return err
	}
	s.AddMeasurement(reading, time.Now())
	return nil
}

// AddMeasurement adds a new measurement to the Sensor
func (s *sensor) AddMeasurement(value float64, timestamp time.Time) {
	s.measurements = append(s.measurements, measurement{s, number(value), timestamp.UTC()})
}
