package main

import (
	"fmt"
)

func removeDuplicates(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		seen := make(map[int]bool)
		for i := range input {
			if !seen[i] {
				seen[i] = true
				output <- i
			}
		}
	}()
	return output
}

func filterEven(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for i := range input {
			if i%2 == 0 {
				output <- i
			}
		}
	}()
	return output
}

func sumBatches(input <-chan int, batchSize int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		sum := 0
		count := 0
		for i := range input {
			sum += i
			count++
			if count == batchSize {
				output <- sum
				sum = 0
				count = 0
			}
		}
		if count > 0 {
			output <- sum
		}
	}()
	return output
}

func main() {
	input := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			input <- i
			input <- i // дубликаты
		}
		close(input)
	}()

	// Создаем pipeline
	noDups := removeDuplicates(input)
	evenOnly := filterEven(noDups)
	batchSums := sumBatches(evenOnly, 5)

	// Выводим результаты
	for sum := range batchSums {
		fmt.Println(sum)
	}
}
