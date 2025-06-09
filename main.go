package main

import (
    "fmt"
    
)

type Vehicle struct {
    Brand string
    Model string
    Year  int
    Color string
}

type VehicleError struct {
    VehicleType string
    Err         error
}

func (ve VehicleError) Error() string {
    return fmt.Sprintf("Vehicle Error: %s - %v", ve.VehicleType, ve.Err)
}

type VehicleInterface interface {
    Start()
    Stop()
    Steer()
}

type Car struct {
    Vehicle
    NumDoors   int
    EngineType string
}

type boat struct {
    Vehicle
    Lenght         int
    PropulsionType string
}

type Motorcycle struct {
    Vehicle
    Wheels      int
    HasSideCar  bool
}

// --- Car methods  ---
func (c Car) Start() {
    fmt.Printf("Starting the %s %s with %d doors\n", c.Color, c.Brand, c.NumDoors)
}

func (c Car) Stop() {
    fmt.Printf("Stopping the %s %s with %d doors\n", c.Color, c.Brand, c.NumDoors)
}

func (c Car) Steer() {
    fmt.Printf("Steering the %s %s with %d doors\n", c.Color, c.Brand, c.NumDoors)
}

// --- boat methods ---
func (b boat) Start() {
    fmt.Printf("Starting the %s %s of length %d with %s propulsion\n",
        b.Color, b.Brand, b.Lenght, b.PropulsionType)
}

func (b boat) Stop() {
    fmt.Printf("Stopping the %s %s of length %d with %s propulsion\n",
        b.Color, b.Brand, b.Lenght, b.PropulsionType)
}

func (b boat) Steer() {
    fmt.Printf("Steering the %s %s of length %d with %s propulsion\n",
        b.Color, b.Brand, b.Lenght, b.PropulsionType)
}

// --- Motorcycle methods ---
func (m Motorcycle) Start() {
    fmt.Printf("Starting the %s %s with %d wheels and sidecar: %t\n",
        m.Color, m.Brand, m.Wheels, m.HasSideCar)
}

func (m Motorcycle) Stop() {
    fmt.Printf("Stopping the %s %s with %d wheels and sidecar: %t\n",
        m.Color, m.Brand, m.Wheels, m.HasSideCar)
}

func (m Motorcycle) Steer() {
    fmt.Printf("Steering the %s %s with %d wheels and sidecar: %t\n",
        m.Color, m.Brand, m.Wheels, m.HasSideCar)
}



func main() {

    c := Car{Vehicle: Vehicle{Brand: "Toyota", Model: "Corolla", Year: 2020, Color: "Blue"}, NumDoors: 4}
    b := boat{Vehicle: Vehicle{Brand: "Yamaha", Model: "WaveRunner", Year: 2019, Color: "White"}, Lenght: 8, PropulsionType: "jet"}
    m := Motorcycle{Vehicle: Vehicle{Brand: "Harley", Model: "Sportster", Year: 2021, Color: "Black"}, Wheels: 2, HasSideCar: false}

    c.Start(); c.Steer(); c.Stop()
    b.Start(); b.Steer(); b.Stop()
    m.Start(); m.Steer(); m.Stop()
}
