package vehicle

// VehicleInterface defines the common actions.
type VehicleInterface interface {
	Start() error
	Stop() error
	Steer(direction string) error
}
