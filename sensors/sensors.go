package sensors

type HardwareDeviceI interface {
	Read(string) (string, error)
	ReadFloat(string) (float64, error)
	Name() (string, error)
}

type SensorDevice struct {
	device HardwareDeviceI
}
