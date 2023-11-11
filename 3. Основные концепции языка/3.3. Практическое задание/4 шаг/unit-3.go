package main

import (
	"fmt"
)

// mostCommonRune Функция для поиска наиболее часто встречающегося символа в строке рун
func mostCommonRune(s string) rune {
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}
	var maxRune rune
	maxCount := 0
	for r, c := range counts {
		if c > maxCount {
			maxCount = c
			maxRune = r
		}
	}
	return maxRune
}

func main() {
	fmt.Println(string(mostCommonRune("Привет, мир!"))) // р
}
