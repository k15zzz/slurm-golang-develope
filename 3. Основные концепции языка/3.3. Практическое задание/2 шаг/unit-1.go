package main

import "fmt"

// reverseSlice Реверс слайса
func reverseSlice(s []int) []int {
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
	return s
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(reverseSlice(s)) // [5 4 3 2 1]
}
