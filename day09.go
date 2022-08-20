package main

import (
	"mathutil"
	"os"
	"panic"
	"readers"
	"sort"
	"strconv"
	"strings"
)

func Day09Part1() int {
	file, err := os.Open("assets/day09.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	heightMap := GetHeightMap(lines)

	riskLevel := 0

	for _, point := range GetLowPoints(heightMap) {
		riskLevel += heightMap[point] + 1
	}

	return riskLevel
}

func Day09Part2() int {
	file, err := os.Open("assets/day09.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	heightMap := GetHeightMap(lines)
	basins := make([]int, len(heightMap))

	for _, point := range GetLowPoints(heightMap) {
		basin := make(map[mathutil.Point]struct{})
		queue := []mathutil.Point{point}

		for len(queue) > 0 {
			currentPoint := queue[0]
			queue = queue[1:]

			if _, isInBasin := basin[currentPoint]; !isInBasin {
				basin[currentPoint] = struct{}{}

				for _, neighbour := range currentPoint.GetNeighbours() {
					if height, isOnMap := heightMap[neighbour]; isOnMap && height < 9 {
						queue = append(queue, neighbour)
					}
				}
			}
		}

		basins = append(basins, len(basin))
	}

	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})

	return basins[0] * basins[1] * basins[2]
}

func GetLowPoints(heightMap map[mathutil.Point]int) []mathutil.Point {
	var lowPoints []mathutil.Point
	for point, height := range heightMap {
		neighbours := point.GetNeighbours()
		foundLower := false
		for _, neighbour := range neighbours {
			neighbourHeight, exists := heightMap[neighbour]

			if exists && neighbourHeight <= height {
				foundLower = true
				break
			}
		}

		if !foundLower {
			lowPoints = append(lowPoints, point)
		}
	}

	return lowPoints
}

func GetHeightMap(lines []string) map[mathutil.Point]int {
	heightMap := make(map[mathutil.Point]int)

	for y, line := range lines {
		heightStrings := strings.Split(line, "")
		for x, heightString := range heightStrings {
			height, err := strconv.Atoi(heightString)
			panic.Check(err)
			heightMap[mathutil.Point{X: x, Y: y}] = height
		}
	}

	return heightMap
}
