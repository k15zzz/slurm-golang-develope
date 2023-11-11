package main

import (
	"fmt"
)

// reverseRunes Функция для переворачивания строки рун
func reverseRunes(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(reverseRunes("Привет")) // тевирП
}
