package sensebox

type SenseBoxSensor struct {
	SensorType string `json:"sensorType"`
	ID         string `json:"_id"`
}

type SenseBox struct {
	ID      string
	Sensors []SenseBoxSensor `json:"sensors"`
}
