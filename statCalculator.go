package main

import (
	"math"
)

func statCalculator(stats map[string]int, level int) map[string]int {
	curStats := map[string]int{}
	for name, val := range stats {
		if name == "hp" {
			curStats[name] = int(math.Ceil((2*float64(val)*float64(level))/100 + float64(level) + 10))
		} else {
			curStats[name] = int(math.Ceil((2*float64(val)*float64(level))/100 + 5))
		}
	}
	return curStats
}