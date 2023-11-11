package main

import (
	"fmt"
	"strings"
	"unicode"
)

// removeSpaces Функция для удаления всех пробелов из строки рун
func removeSpaces(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

func main() {
	fmt.Println(removeSpaces("Привет, мир!")) // тевирП
}
