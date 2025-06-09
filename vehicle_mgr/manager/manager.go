package manager

import (
	"fmt"
	"sync"

	"github.com/showbart/vehicle_mgr/vehicle"
)

// Manager keeps track of registered vehicles.
type Manager struct {
	mu       sync.Mutex
	vehicles []vehicle.VehicleInterface
}

// Add a new vehicle to the system.
func (m *Manager) Add(v vehicle.VehicleInterface) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.vehicles = append(m.vehicles, v)
}

// Start a vehicle by index; returns error if any.
func (m *Manager) Start(idx int) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if idx < 0 || idx >= len(m.vehicles) {
		return fmt.Errorf("no vehicle at index %d", idx)
	}
	return m.vehicles[idx].Start()
}

// Stop a vehicle by index.
func (m *Manager) Stop(idx int) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if idx < 0 || idx >= len(m.vehicles) {
		return fmt.Errorf("no vehicle at index %d", idx)
	}
	return m.vehicles[idx].Stop()
}
