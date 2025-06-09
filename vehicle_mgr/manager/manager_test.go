package manager

import (
	"testing"

	"github.com/showbart/vehicle_mgr/vehicle"
)

type stubV struct { vehicle.Vehicle }
func (s *stubV) Start() error   { return nil }
func (s *stubV) Stop() error    { return nil }
func (s *stubV) Steer(_ string) error { return nil }

func TestAddAndStart(t *testing.T) {
	m := &Manager{}
	v := &stubV{}
	m.Add(v)
	if err := m.Start(0); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
