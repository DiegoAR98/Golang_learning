package assignment1

import "fmt"

func CreateMap() {
	mappy := make(map[int]string)
	{
		mappy[1] = "One"
		mappy[2] = "Two"
		mappy[3] = "Three"
	}
	fmt.Println(mappy)
}

func DeleteMap(mappyA map[int]string, key int) {
	for key, _ := range mappyA {

		delete(mappyA, key)
		fmt.Println("After deletion:", mappyA)
	}

}
