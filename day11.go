package main

import (
	"mathutil"
	"os"
	"panic"
	"readers"
)

func Day11Part1() int {
	file, err := os.Open("assets/day11.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	energyMap := mathutil.GetIntMap(lines)

	flashes := 0

	for i := 0; i < 100; i++ {
		for point := range energyMap {
			energyMap[point]++
		}

		for IsAnyReady(energyMap) {
			for point, value := range energyMap {
				if value > 9 && value < 10000 {
					for _, adjacentPoint := range point.GetAdjacent() {
						if _, isOnMap := energyMap[adjacentPoint]; isOnMap {
							energyMap[adjacentPoint]++
						}
					}
					energyMap[point] = 10000
					flashes++
				}
			}
		}

		for point, value := range energyMap {
			if value > 9 {
				energyMap[point] = 0
			}
		}
	}

	return flashes
}

func Day11Part2() int {
	file, err := os.Open("assets/day11.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	energyMap := mathutil.GetIntMap(lines)

	steps := 0
	for true {
		flashes := 0
		steps++
		for point := range energyMap {
			energyMap[point]++
		}

		for IsAnyReady(energyMap) {
			for point, value := range energyMap {
				if value > 9 && value < 10000 {
					for _, adjacentPoint := range point.GetAdjacent() {
						if _, isOnMap := energyMap[adjacentPoint]; isOnMap {
							energyMap[adjacentPoint]++
						}
					}
					energyMap[point] = 10000
					flashes++
				}
			}
		}

		if flashes == 100 {
			return steps
		}

		for point, value := range energyMap {
			if value > 9 {
				energyMap[point] = 0
			}
		}
	}

	return -1
}

func IsAnyReady(energyMap map[mathutil.Point]int) bool {
	for _, value := range energyMap {
		if value > 9 && value < 10000 {
			return true
		}
	}

	return false
}
