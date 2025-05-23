package slices

import "fmt"

func LearnSlices() {
	var slice []int
	slice = append(slice, 0)
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	fmt.Println(slice)

	sliceWithSize := make([]int, 7)
fmt.Print(sliceWithSize)
}