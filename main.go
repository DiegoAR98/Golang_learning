package main

import (
	"fmt"
)

// type person struct {
// 	age int
// 	name string
// }

// func (p *person) changeUser(){
// 	p.age = 20
// 	p.name = "John Doe"
// }

// type child struct {
// 	age  int
// 	name string
// 	adult.Adult
// }

type Car string
type Motorcycle string

type Drive interface {
	Steer()
}

func (c Car) Steer() {
	fmt.Println("Car is steering")
}

func (m Motorcycle) Steer() {
	fmt.Println("Motorcycle is steering")
}

func canSteer(d Drive) {
	d.Steer()
}

func main() {
	canSteer(Car("Toyota"))
	canSteer(Motorcycle("Harley Davidson"))
}

type error interface {
	Error() string
}
