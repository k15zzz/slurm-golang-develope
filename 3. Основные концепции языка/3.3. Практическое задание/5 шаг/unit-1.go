package main

import (
	"fmt"
)

// swap Функция, которая меняет местами две переменных через указатель
func swap[T comparable](a, b *T) {
	*a, *b = *b, *a
}

func main() {
	a := "Hello"
	b := "Golang"
	swap(&a, &b)
	fmt.Println(a) // Golang
}
