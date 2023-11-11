package main

import (
	"fmt"
	"math"
)

// compareCoordinates Сравнение координат
func compareCoordinates(lat1, lon1, lat2, lon2 float64) bool {
	const epsilon = 0.00001
	return math.Abs(lat1-lat2) < epsilon && math.Abs(lon1-lon2) < epsilon
}

func main() {
	fmt.Println(compareCoordinates(36.12, -86.67, 36.12, -86.67))  // true
	fmt.Println(compareCoordinates(36.12, -86.67, 33.94, -118.40)) // false
}
