package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int, requests <-chan int, results chan<- int) {
	defer wg.Done()
	for r := range requests {
		fmt.Println("Worker", id, "finished work on request", r)
		time.Sleep(100 * time.Millisecond)
		results <- r * 4
	}
}

func main() {
	fmt.Println("Time started", time.Now())
	jobs := make(chan int, 100)
	results := make(chan int, 1000)

	// Start worker threads.
	// Start: 10:54:33.754024
	// End: 10:54:40.188991 // For 16 threads. Time: <7s
	// Start: 10:55:28.736481
	// End: 10:57:12.020337 // For 1 thread. Time: 1m 44s
	for w := 1; w <= 16; w++ {
		wg.Add(1)
		go worker(w, jobs, results)
	}

	// Continuously put 1000 requests to the jobs channel.
	// Even though the jobs only has buffer size of 100. This is the sliding throttling window.
	for r := 1; r <= 1000; r++ {
		jobs <- r
	}
	close(jobs)

	for a := 1; a <= 1000; a++ {
		<-results
	}
	wg.Wait()
	fmt.Println("Time finished", time.Now())
}
