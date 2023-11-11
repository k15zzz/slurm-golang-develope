package main

import (
	"fmt"
	"regexp"
	"strings"
)

// findMostWord Поиск наиболее часто встречающегося слова в строке
func findMostWord(str string) (string, int) {
	var re = regexp.MustCompile(`[[:punct:]]`)
	words := strings.Fields(strings.ToLower(re.ReplaceAllString(str, "")))
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++
	}
	var mostFrequentWord string
	var mostFrequentCount int
	for word, count := range wordCounts {
		if count > mostFrequentCount {
			mostFrequentWord = word
			mostFrequentCount += count
		}
	}
	return mostFrequentWord, mostFrequentCount
}

// findMostFrequentWordBySlice Поиск наиболее часто встречающегося слова в слайсе строк
func findMostFrequentWordBySlice(sliceStr []string) string {
	wordCounts := make(map[string]int)
	for _, value := range sliceStr {
		mostWord, mostCount := findMostWord(value)
		wordCounts[mostWord] += mostCount
	}
	var mostFrequentWord string
	var mostFrequentCount int
	for word, count := range wordCounts {
		if count > mostFrequentCount {
			mostFrequentWord = word
			mostFrequentCount += count
		}
	}
	return mostFrequentWord
}

func main() {
	m := []string{
		"Ехали, ехали и наконец - приехали",
		"Ехали, ехали, ехали и наконец - приехали приехали приехали",
		"Ехали, ехали и наконец - приехали",
	}
	fmt.Println(findMostFrequentWordBySlice(m)) //ехали
}
