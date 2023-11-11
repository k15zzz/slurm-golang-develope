package main

import (
	"fmt"
	"regexp"
	"strings"
)

// findMostFrequentWord Поиск наиболее часто встречающегося слова в строке
func findMostFrequentWord(str string) (string, int) {
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
			mostFrequentCount = count
		}
	}
	return mostFrequentWord, mostFrequentCount
}

func main() {
	fmt.Println(findMostFrequentWord("Ехали, ехали и наконец - приехали")) // ехали
}
