package main

import "fmt"

func mapIntersection(m map[string]int) map[string]int {
	result := make(map[string]int)
	for key := range m {
		fmt.Println(key)
	}
	return result
}

func main() {
	mapIntersection(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	})
}
