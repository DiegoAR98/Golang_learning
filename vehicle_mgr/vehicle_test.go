package main

import (
	"testing"
	"time"

	"github.com/yourname/vehicle_mgr/vehicle"
)

func TestNewError(t *testing.T) {
	err := vehicle.NewError("start-fail", "engine cold")
	ve, ok := err.(*vehicle.VehicleError)
	if !ok {
		t.Fatal("expected VehicleError")
	}
	if ve.What != "start-fail" {
		t.Errorf("want What=start-fail, got %q", ve.What)
	}
	// you can also check time roughly
	if time.Since(ve.When) > time.Second {
		t.Error("When not set correctly")
	}
}
