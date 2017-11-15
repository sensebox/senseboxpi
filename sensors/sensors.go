package sensors

// HardwareDeviceI is a interface for all sorts of hardware sensors
type HardwareDeviceI interface {
	Read(string) (string, error)
	ReadFloat(string) (float64, error)
	Name() (string, error)
}

// SensorDevice wraps a device implementing the HardwareDeviceI interface
type SensorDevice struct {
	device HardwareDeviceI
}
