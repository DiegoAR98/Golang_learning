package main

import (
	"fmt"
)

var name struct {
	value int
	num  float32
	str string
}

func main() {
	fmt.Print(name)
}
