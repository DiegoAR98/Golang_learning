package maps

import "fmt"

func LearnMaps() {
	mappy := make(map[int]string)
	mappy[1] = "One"
	mappy[2] = "Two"
	mappy[3] = "Three"
	fmt.Println(mappy)
}