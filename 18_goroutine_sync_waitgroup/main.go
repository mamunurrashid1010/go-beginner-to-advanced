package main

import (
	"fmt"
	"sync"
	"time"
)

/**
	Description :
	Goroutine : A lightweight thread managed by the Go runtime, used for concurrency.
	sync.WaitGroup : A synchronization primitive that waits for a collection of goroutines to finish executing.
**/

func worker(id int, wg *sync.WaitGroup) {
	// Decrement the WaitGroup counter when the function completes
	defer wg.Done()

	fmt.Printf("Worker %d: starting\n", id)

	// Simulate work by sleeping for 2 seconds
	time.Sleep(2 * time.Second)

	fmt.Printf("Worker %d: finished\n", id)
}

func main() {
	// Define a waitgroup
	var wg sync.WaitGroup

	// Lunch 3 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment counter for each goroutine
		go worker(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All workers finished!")
}
