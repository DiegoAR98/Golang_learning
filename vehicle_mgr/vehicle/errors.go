package vehicle

import "fmt"

// VehicleError holds context when something goes wrong.
type VehicleError struct {
	When   time.Time
	What   string
	Detail string
}

func (e *VehicleError) Error() string {
	return fmt.Sprintf("%s: %s (%s)", e.When.Format(time.RFC3339), e.What, e.Detail)
}

// helper to create one
func NewError(what, detail string) error {
	return &VehicleError{
		When:   time.Now(),
		What:   what,
		Detail: detail,
	}
}
