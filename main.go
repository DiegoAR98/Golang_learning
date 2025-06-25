package main

import "fmt"

type Num interface {
    int | int8 | int16 | int32 | int64 | float32 | float64
}

func Add[T comparable, V Num](m map[T]V) V {
    var sum V 
    for _, v := range m {
        sum += v
    }
    return sum
}

func main() {  
    ints := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
}

 floats := map[string]float64{
        "x": 1.1,
}

result := Add(ints)
    fmt.Println("Sum of ints:", result) 

    resultFloat := Add(floats)
    fmt.Println("Sum of floats:", resultFloat)
    
}