package main

import (
	"fmt"
)

// bubbleSort Пузырьковая сортировка слайса
func bubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func main() {
	slice := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted slice: ", slice)
	bubbleSort(slice)
	fmt.Println("Sorted slice: ", slice)
}
