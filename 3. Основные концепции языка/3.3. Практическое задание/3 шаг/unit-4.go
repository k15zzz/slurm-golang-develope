package main

import (
	"fmt"
)

// findingIntersectionMaps Поиск пересечения двух map
func findingIntersectionMaps(m1 map[string]string, m2 map[string]string) map[string]string {
	m := make(map[string]string)
	for key1, _ := range m1 {
		value2, ok := m2[key1]
		if ok {
			m[key1] = value2
		}
	}
	return m
}

func main() {
	m1 := map[string]string{
		"q": "key",
		"w": "key",
		"e": "key",
		"r": "key",
	}
	m2 := map[string]string{
		"t": "key",
		"w": "key",
		"y": "key",
		"r": "key",
	}
	fmt.Println(findingIntersectionMaps(m1, m2)) // ехали
}
