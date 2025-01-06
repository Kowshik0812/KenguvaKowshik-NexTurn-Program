package main

import (
	"fmt"
	"strings"
)

type City struct {
	Name     string
	AvgTemp  float64
	Rainfall float64
}

func main() {
	cities := []City{
		{"Delhi", 32.4, 800.0},
		{"Visakapatanam", 34.7, 1200.0},
		{"Hyderabad", 35.2, 1300.5},
		{"Banglore", 29.3, 1500.0},
		{"Mumbai", 31.1, 507.4},
	}

	highestTempCity, lowestTempCity := findTemperatureExtremes(cities)
	fmt.Printf("City with the highest average temperature: %s (%.2f°C)\n", highestTempCity.Name, highestTempCity.AvgTemp)
	fmt.Printf("City with the lowest average temperature: %s (%.2f°C)\n", lowestTempCity.Name, lowestTempCity.AvgTemp)

	averageRainfall := calculateAverageRainfall(cities)
	fmt.Printf("Average rainfall across all cities: %.2f mm\n", averageRainfall)

	var threshold float64
	fmt.Print("Enter rainfall threshold (mm): ")
	fmt.Scanln(&threshold)
	filterCitiesByRainfall(cities, threshold)

	var cityName string
	fmt.Print("Enter city name to search: ")
	fmt.Scanln(&cityName)
	searchCityByName(cities, cityName)
}

func findTemperatureExtremes(cities []City) (City, City) {
	highest := cities[0]
	lowest := cities[0]
	for _, city := range cities[1:] {
		if city.AvgTemp > highest.AvgTemp {
			highest = city
		}
		if city.AvgTemp < lowest.AvgTemp {
			lowest = city
		}
	}
	return highest, lowest
}

func calculateAverageRainfall(cities []City) float64 {
	totalRainfall := 0.0
	for _, city := range cities {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cities))
}

func filterCitiesByRainfall(cities []City, threshold float64) {
	fmt.Printf("Cities with rainfall above %.2f mm:\n", threshold)
	for _, city := range cities {
		if city.Rainfall > threshold {
			fmt.Printf("- %s (%.2f mm)\n", city.Name, city.Rainfall)
		}
	}
}

func searchCityByName(cities []City, name string) {
	found := false
	name = strings.ToLower(name)
	for _, city := range cities {
		if strings.ToLower(city.Name) == name {
			fmt.Printf("Data for %s: Avg Temp: %.2f°C, Rainfall: %.2f mm\n", city.Name, city.AvgTemp, city.Rainfall)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("City '%s' not found in the data.\n", name)
	}
}
