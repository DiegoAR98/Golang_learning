package arrays

import "fmt"

func Learn() {
	var array [5]int
	array[0] = 0
	array[1] = 1
	array[2] = 2
	array[3] = 3
	array[4] = 4

	// literal := [5]int{0, 1, 2, 3, 4}

	// fmt.Println(array)
	// fmt.Println(literal)

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	for _, value := range array {
		fmt.Println(value)
	}
}
