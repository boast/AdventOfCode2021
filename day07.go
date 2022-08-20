package main

import (
	"collections"
	"math"
	"mathutil"
	"os"
	"panic"
	"readers"
	"strings"
)

func Day07Part1() int {
	file, err := os.Open("assets/day07.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	positionStrings := strings.Split(lines[0], ",")
	positions := collections.StringSliceToIntSlice(positionStrings)

	median := mathutil.MedianInt(positions)

	fuelRequired := collections.MapInt(positions, func(i int) int {
		distance := math.Abs(float64(i) - median)
		return int(distance)
	})

	return collections.SumInt(fuelRequired)
}

func Day07Part2() int {
	file, err := os.Open("assets/day07.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	positionStrings := strings.Split(lines[0], ",")
	positions := collections.StringSliceToIntSlice(positionStrings)

	average := int(mathutil.AverageInt(positions))

	fuelRequired := collections.MapInt(positions, func(i int) int {
		distance := int(math.Abs(float64(i) - float64(average)))
		return mathutil.SumUpTo(distance)
	})

	return collections.SumInt(fuelRequired)
}
