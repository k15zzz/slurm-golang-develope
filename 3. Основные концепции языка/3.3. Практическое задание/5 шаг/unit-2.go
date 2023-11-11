package main

import (
	"fmt"
)

// copySlice Функция, которая создает копию слайса
func copySlice[T comparable](slice []T) []T {
	newSlice := make([]T, len(slice))
	for i, v := range slice {
		newSlice[i] = v
	}
	return newSlice
}

func main() {
	slice := []string{
		"q",
		"w",
		"r",
	}
	fmt.Println(copySlice(slice)) // тевирП
}
