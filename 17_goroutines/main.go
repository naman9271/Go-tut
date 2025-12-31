package main

import (
	"fmt"
	"sync"
	// "time"
)

// Goroutines in Go
// A goroutine is a lightweight thread managed by the Go runtime.
// You can start a goroutine by using the 'go' keyword followed by a function call.

// WaitGroup can also be used for better synchronization, but here we use Sleep for simplicity.
// Add(n int)   // increase the counter by n
// Done()       // decrease the counter by 1
// Wait()       // block until counter becomes 0

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine is running", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
	// time.Sleep(1 * time.Second) // Wait for goroutines to finish
	fmt.Println("Main function is done")
}
