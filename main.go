package main

import (
	"fmt"
	"sync"
	"time"


)

func printOne(wg *sync.WaitGroup){
    defer wg.Done()
    
    for i := 0; i < 50; i++ {
        fmt.Println("1")
        time.Sleep(10 * time.Millisecond)
}
}

func printTwo(wg *sync.WaitGroup){
    defer wg.Done()
    
    for i := 0; i < 50; i++ {
        fmt.Println("2")
        time.Sleep(10 * time.Millisecond)
}
}

func sum(numbers []int, resultChan chan <- int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    
    resultChan <- sum // Send the result to the channel
}


func main() {
    var wg sync.WaitGroup

    resultChan := make(chan int)  //<- Create a channel to receive results
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    numGoRoutines := 4

    sliceSize := len(numbers) / numGoRoutines

    for i := 0; i < numGoRoutines; i++ {
        start := i * sliceSize
        end := (i + 1) * sliceSize

         if i == numGoRoutines-1 {
        end = len(numbers) // Handle the last slice to include any remaining elements
    }

    wg.Add(1) // Add a goroutine to the wait group
    go sum(numbers[start:end], resultChan, &wg)

    }

    var collectWg sync.WaitGroup
    collectWg.Add(1)

    go func() {
        defer collectWg.Done()
        totalSum := 0
        for i := 0; i < numGoRoutines; i++ {
            partialSum := <-resultChan // Receive the result from the channel
            totalSum += partialSum
        }
        fmt.Println("Total Sum:", totalSum)
    }()


   

    // wg.Add(2) // Add two goroutines to the wait group
    // go printOne(&wg)
    // go printTwo(&wg)

    wg.Wait()
    close(resultChan) // Close the channel after all goroutines are done
    collectWg.Wait() // Wait for the collection goroutine to finish
}
