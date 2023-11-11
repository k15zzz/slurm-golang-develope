package main

import "fmt"

// removeDuplicates Удаление дублей из слайса
func removeDuplicates[T comparable](s []T) []T {
	seen := make(map[T]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}

func main() {
	s := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	fmt.Println(removeDuplicates(s)) // [1 2 3 4]

	a := []string{"a", "a", "a", "a", "a", "a", "b", "c", "b", "b", "b", "b", "b", "b", "d"}
	fmt.Println(removeDuplicates(a)) // [a d b c]

}
