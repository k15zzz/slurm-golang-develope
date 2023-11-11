package main

import "fmt"

// factorialRecursive Рекурсивный подход нахождения факториала
func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

// factorialIterative Итеративный подход нахождения факториала
func factorialIterative(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Println(factorialRecursive(5)) // 120
	fmt.Println(factorialIterative(5)) // 120
}
