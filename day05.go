package main

import (
	"mathutil"
	"os"
	"panic"
	"readers"
	"strconv"
	"strings"
)

func Day05Part1() int {
	file, err := os.Open("assets/day05.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	pointVectors := ParsePointVectors(lines)

	depthMap := make(map[mathutil.Point]int)

	for _, vector := range pointVectors {
		if vector.a.X != vector.b.X && vector.a.Y != vector.b.Y {
			continue
		}

		if vector.a.X == vector.b.X {
			minY := vector.a.Y
			maxY := vector.b.Y

			if minY > maxY {
				minY = vector.b.Y
				maxY = vector.a.Y
			}

			for y := minY; y <= maxY; y++ {
				point := mathutil.Point{X: vector.a.X, Y: y}
				depthMap[point]++
			}
		} else {
			minX := vector.a.X
			maxX := vector.b.X

			if minX > maxX {
				minX = vector.b.X
				maxX = vector.a.X
			}

			for x := minX; x <= maxX; x++ {
				point := mathutil.Point{X: x, Y: vector.a.Y}
				depthMap[point]++
			}
		}

	}

	count := 0

	for _, i := range depthMap {
		if i >= 2 {
			count++
		}
	}

	return count
}

func Day05Part2() int {
	file, err := os.Open("assets/day05.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	pointVectors := ParsePointVectors(lines)

	depthMap := make(map[mathutil.Point]int)

	for _, vector := range pointVectors {
		if vector.a.X != vector.b.X && vector.a.Y != vector.b.Y {
			var slope mathutil.Point

			if vector.a.X < vector.b.X && vector.a.Y < vector.b.Y {
				slope = mathutil.Point{X: 1, Y: 1}
			} else if vector.a.X < vector.b.X && vector.a.Y > vector.b.Y {
				slope = mathutil.Point{X: 1, Y: -1}
			} else if vector.a.X > vector.b.X && vector.a.Y < vector.b.Y {
				slope = mathutil.Point{X: -1, Y: 1}
			} else {
				slope = mathutil.Point{X: -1, Y: -1}
			}

			currentPoint := vector.a

			depthMap[currentPoint]++

			for currentPoint != vector.b {
				currentPoint.Add(slope)
				depthMap[currentPoint]++
			}

			continue
		}

		if vector.a.X == vector.b.X {
			minY := vector.a.Y
			maxY := vector.b.Y

			if minY > maxY {
				minY = vector.b.Y
				maxY = vector.a.Y
			}

			for y := minY; y <= maxY; y++ {
				point := mathutil.Point{X: vector.a.X, Y: y}
				depthMap[point]++
			}
		} else {
			minX := vector.a.X
			maxX := vector.b.X

			if minX > maxX {
				minX = vector.b.X
				maxX = vector.a.X
			}

			for x := minX; x <= maxX; x++ {
				point := mathutil.Point{X: x, Y: vector.a.Y}
				depthMap[point]++
			}
		}

	}

	count := 0

	for _, i := range depthMap {
		if i >= 2 {
			count++
		}
	}

	return count
}

type PointVector struct {
	a mathutil.Point
	b mathutil.Point
}

func ParsePointVectors(lines []string) []PointVector {
	var vectors []PointVector

	for _, line := range lines {
		var coordinates []int

		for _, pointValues := range strings.Split(line, " -> ") {
			for _, coordinate := range strings.Split(pointValues, ",") {
				coordinateInt, err := strconv.Atoi(coordinate)
				panic.Check(err)
				coordinates = append(coordinates, coordinateInt)
			}
		}

		vectors = append(vectors, PointVector{
			a: mathutil.Point{X: coordinates[0], Y: coordinates[1]},
			b: mathutil.Point{X: coordinates[2], Y: coordinates[3]},
		})
	}

	return vectors
}
