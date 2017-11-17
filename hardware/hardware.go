package hardware

// DeviceI is the interface devices should implement
type DeviceI interface {
	Read(string) (string, error)
	ReadFloat(string) (float64, error)
	Name() (string, error)
}
