package main

import (
	"fmt"
	"math"
)

// haversine Расстояние между двумя точками на поверхности Земли
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371 // радиус Земли в километрах

	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180

	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := earthRadius * c
	return distance
}

func main() {
	fmt.Println(haversine(36.12, -86.67, 33.94, -118.40)) // Расстояние от Нашвилла, Теннесси до Лос-Анджелеса, Калифорния
}
