package main

import (
	"fmt"
	"math"
	"sort"
)

func groupTemperatures(temperatures []float64, step float64) map[int][]float64 {
	groups := make(map[int][]float64)

	roundDown := func(value, step float64) float64 {
		return math.Floor(value/step) * step
	}

	sort.Float64s(temperatures)

	for _, temp := range temperatures {
		groupKey := int(roundDown(temp, step))
		groups[groupKey] = append(groups[groupKey], temp)
	}

	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10.0

	groups := groupTemperatures(temperatures, step)

	for key, values := range groups {
		fmt.Printf("%d: %v\n", key, values)
	}
}
