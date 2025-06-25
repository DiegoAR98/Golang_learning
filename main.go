package main

import (
    "fmt"
    "sync"
)


func increment(counter *int, mu *sync.Mutex) {
    mu.Lock()
    defer mu.Unlock()
    *counter++
}

func main() {  
    numGoRoutines := 100
    var wg sync.WaitGroup
    var counter int
    var mu sync.Mutex

    for i := 0; i < numGoRoutines; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            increment(&counter, &mu)
        }()
    }
    wg.Wait()
    fmt.Println("Final counter value:", counter)
}
