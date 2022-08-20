package main

import (
	"mathutil"
	"os"
	"panic"
	"readers"
	"strconv"
	"strings"
)

func Day13Part1() int {
	file, err := os.Open("assets/day13.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	display := ProcessLines(lines, true)

	return len(display)
}

func Day13Part2() int {
	file, err := os.Open("assets/day13.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	display := ProcessLines(lines, false)

	display.Print()

	return len(display)
}

func ProcessLines(lines []string, breakOnFirstFold bool) mathutil.PointMap {
	display := make(mathutil.PointMap)

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.Contains(line, ",") {
			coordinates := strings.Split(line, ",")

			x, err := strconv.Atoi(coordinates[0])
			panic.Check(err)
			y, err := strconv.Atoi(coordinates[1])
			panic.Check(err)

			display[mathutil.Point{X: x, Y: y}] = struct{}{}
			continue
		} else if strings.Contains(line, "fold along x=") {
			coordinate := strings.Replace(line, "fold along x=", "", 1)
			x, err := strconv.Atoi(coordinate)
			panic.Check(err)

			display.FoldX(x)

			if breakOnFirstFold {
				break
			} else {
				continue
			}
		} else if strings.Contains(line, "fold along y=") {
			coordinate := strings.Replace(line, "fold along y=", "", 1)
			y, err := strconv.Atoi(coordinate)
			panic.Check(err)

			display.FoldY(y)

			if breakOnFirstFold {
				break
			} else {
				continue
			}
		}
	}

	return display
}
