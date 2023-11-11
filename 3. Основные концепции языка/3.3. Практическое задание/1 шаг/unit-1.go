package main

import (
	"fmt"
	"math"
)

// addConstant Переполнение int32
func addConstant(i int32) int32 {
	return math.MaxInt32 - i
}

func main() {
	//fmt.Println(addConstant(2)) // Это приведет к переполнению
	fmt.Println(addConstant(999999999))
}
