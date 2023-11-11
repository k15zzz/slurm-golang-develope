package main

import (
	"fmt"
)

// addElementToSlice добавление нового элемента в слайс
func addElementToSlice[T comparable](slice []T, element T) []T {
	newSlice := make([]T, len(slice)+1)
	copy(newSlice, slice)
	newSlice[len(slice)] = element
	return newSlice
}

func main() {
	slice := []int{1, 2, 3, 4}
	fmt.Println(addElementToSlice(slice, 5)) // [1 2 3 4 5]
}
