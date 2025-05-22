package main

import (
	"fmt"
)

func Hello() {
	fmt.Println("Hello, World!")
}

func Add(x int, y int) (bool, int) {
	if x+y != 0 {
		return true, x + y
	} else {
		return false, x + y
	}
}

func main() {
	value := 1
	fmt.Println(value)
	fmt.Print(&value)
}
