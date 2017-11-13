package sensebox

import "github.com/sensebox/senseboxpi/sensors"

type SenseBoxSensor struct {
	SensorType string `json:"sensorType"`
	ID         string `json:"_id"`
}

type SenseBox struct {
	ID      string           `json:"_id"`
	Sensors []SenseBoxSensor `json:"sensors"`
}

type SenseBoxSensorDevice struct {
	Sensor       SenseBoxSensor
	SensorDevice sensors.SensorDevice
}

func NewSenseBoxSensorDevice(senseBoxSensor SenseBoxSensor, sensorDevice sensors.SensorDevice) SenseBoxSensorDevice {
	return SenseBoxSensorDevice{Sensor: senseBoxSensor, SensorDevice: sensorDevice}
}
