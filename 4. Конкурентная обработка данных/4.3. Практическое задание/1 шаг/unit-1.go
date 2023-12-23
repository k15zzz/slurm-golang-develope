package main

import (
	"fmt"
	"sync"
)

func worker(id int, numbers <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for n := range numbers {
		sum += n
	}
	fmt.Printf("Worker %d finished with sum %d\n", id, sum)
	results <- sum
}

func writeNumbers(numNumbers int, numbers chan<- int) {
	for i := 0; i < numNumbers; i++ {
		numbers <- i
	}
	close(numbers)
}

func printSum(results <-chan int) {
	totalSum := 0
	for sum := range results {
		totalSum += sum
	}
	fmt.Printf("Total sum: %d\n", totalSum)
}

func main() {
	const numWorkers = 5
	const numNumbers = 88

	numbers := make(chan int, numNumbers)
	results := make(chan int, numWorkers)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(i, numbers, results, &wg)
	}

	go writeNumbers(numNumbers, numbers)

	wg.Wait()
	close(results)

	printSum(results)
}
