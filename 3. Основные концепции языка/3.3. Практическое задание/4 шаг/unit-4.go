package main

import (
	"fmt"
	"unicode"
)

// lengthWithoutSpaces Функция для поиска длины строки без учета пробелов
func lengthWithoutSpaces(s string) int {
	count := 0
	for _, r := range s {
		if !unicode.IsSpace(r) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(lengthWithoutSpaces("Привет, мир!")) // р
}
