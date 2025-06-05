package main

import (
	"fmt"


	"example.com/course/adult"
)

// type person struct {
// 	age int
// 	name string
// }

// func (p *person) changeUser(){
// 	p.age = 20
// 	p.name = "John Doe"
// }

type child struct {
	age  int
	name string
	adult.Adult
}

func main() {
	// parent := adult.Adult{30, "Bob"}

	// err := parent.SetAge(-1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(parent.GetAge())

	kid := child{
		5, "cris", adult.Adult{}}

		fmt.Println(kid.age)
		kid.Adult.SetAge(10)
		fmt.Println(kid.Adult.GetAge())
}
