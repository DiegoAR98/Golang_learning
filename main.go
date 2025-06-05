package main

import (
	"fmt"
)

type person struct {
	age int
	name string
}

func (p *person) changeUser(){
	p.age = 20
	p.name = "John Doe"
}

func main() {
	var user person = person{30, "Jane Doe"}
	user.changeUser()
	fmt.Println(user)
}
