package main

import (
	"math"
	"mathutil"
	"os"
	"panic"
	"readers"
	"strconv"
	"strings"
)

type TargetArea struct {
	xMin int
	xMax int
	yMin int
	yMax int
}

func Day17Part1() int {
	file, err := os.Open("assets/day17.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	targetArea := GetTargetRange(lines[0])
	yMax := math.MinInt

	for y := -1000; y < 1000; y++ {
		for x := -1000; x < 1000; x++ {
			if DidHit, yMaxForHit := TryHitTarget(x, y, targetArea); DidHit {
				if yMaxForHit > yMax {
					yMax = yMaxForHit
				}
			}
		}
	}

	return yMax
}

func Day17Part2() int {
	file, err := os.Open("assets/day17.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	targetArea := GetTargetRange(lines[0])

	hitCount := 0

	for y := -1000; y < 1000; y++ {
		for x := -1000; x < 1000; x++ {
			if DidHit, _ := TryHitTarget(x, y, targetArea); DidHit {
				hitCount++
			}
		}
	}

	return hitCount
}

func TryHitTarget(xVelocity, yVelocity int, targetArea TargetArea) (bool, int) {
	position := mathutil.Point{}

	yMax := 0

	for position.X <= targetArea.xMax && position.Y >= targetArea.yMin {
		if position.Y > yMax {
			yMax = position.Y
		}
		if position.X >= targetArea.xMin && position.Y <= targetArea.yMax {
			return true, yMax
		}

		position.X += xVelocity
		position.Y += yVelocity

		if xVelocity > 0 {
			xVelocity--
		} else if xVelocity < 0 {
			xVelocity++
		}
		yVelocity--
	}

	return false, yMax
}

func GetTargetRange(line string) TargetArea {
	targetArea := TargetArea{}
	var err error

	targets := strings.Split(strings.Replace(line, "target area: ", "", 1), ", ")

	xRangeStrings := strings.Split(strings.Replace(targets[0], "x=", "", 1), "..")
	yRangeStrings := strings.Split(strings.Replace(targets[1], "y=", "", 1), "..")

	targetArea.xMin, err = strconv.Atoi(xRangeStrings[0])
	panic.Check(err)
	targetArea.xMax, err = strconv.Atoi(xRangeStrings[1])
	panic.Check(err)
	targetArea.yMin, err = strconv.Atoi(yRangeStrings[0])
	panic.Check(err)
	targetArea.yMax, err = strconv.Atoi(yRangeStrings[1])
	panic.Check(err)

	return targetArea
}
