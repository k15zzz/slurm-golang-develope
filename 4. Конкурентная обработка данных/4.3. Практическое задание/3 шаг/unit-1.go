package main

import (
	"fmt"
	//"golang.org/x/sync/semaphore"
	"sync"
	"time"
)

type Semaphore chan struct{}

func NewSemaphore(n int) Semaphore {
	return make(Semaphore, n)
}

func (s Semaphore) Acquire(n int) {
	e := struct{}{}
	for i := 0; i < n; i++ {
		s <- e
	}
}

func (s Semaphore) Release(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

const (
	maxRequests    = 100
	maxConcurrency = 10
)

func main() {
	sem := NewSemaphore(maxConcurrency)

	var wg sync.WaitGroup

	for i := 0; i < maxRequests; i++ {
		wg.Add(1)

		sem.Acquire(1)

		go func(id int) {
			defer wg.Done()
			defer sem.Release(1)
			startTime := time.Now()
			time.Sleep(time.Second)
			fmt.Printf("Request %d completed in %v\n", id, time.Since(startTime))
		}(i)
	}

	wg.Wait()
	fmt.Println("All requests have been processed.")
}
