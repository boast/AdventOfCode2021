package main

import (
	"os"
	"panic"
	"readers"
	"strconv"
	"strings"
)

func Day22Part1() int64 {
	file, err := os.Open("assets/day22.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	cuboids := parseCuboid(lines)

	for i := 0; i < len(cuboids); i++ {
		if cuboids[i].xMin < -50 || cuboids[i].xMax > 50 || cuboids[i].yMin < -50 || cuboids[i].yMax > 50 || cuboids[i].zMin < -50 || cuboids[i].zMax > 50 {
			cuboids = append(cuboids[:i], cuboids[i+1:]...)
			i--
		}
	}

	finalCuboids := make([]*cuboid, 0)

	for _, c := range cuboids {
		for _, finalCuboid := range finalCuboids {
			finalCuboid.storeNegatives(c)
		}

		if c.isOn {
			finalCuboids = append(finalCuboids, c)
		}
	}

	volume := int64(0)

	for _, finalCuboid := range finalCuboids {
		volume += finalCuboid.getVolume()
	}

	return volume
}

func Day22Part2() int64 {
	file, err := os.Open("assets/day22.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	cuboids := parseCuboid(lines)

	finalCuboids := make([]*cuboid, 0)

	for _, c := range cuboids {
		for _, finalCuboid := range finalCuboids {
			finalCuboid.storeNegatives(c)
		}

		if c.isOn {
			finalCuboids = append(finalCuboids, c)
		}
	}

	volume := int64(0)

	for _, finalCuboid := range finalCuboids {
		volume += finalCuboid.getVolume()
	}

	return volume
}

type cuboid struct {
	xMin, xMax int
	yMin, yMax int
	zMin, zMax int
	isOn       bool
	negatives  []*cuboid
}

func newCuboid(xMin, xMax, yMin, yMax, zMin, zMax int) *cuboid {
	return &cuboid{
		xMin:      xMin,
		xMax:      xMax,
		yMin:      yMin,
		yMax:      yMax,
		zMin:      zMin,
		zMax:      zMax,
		negatives: make([]*cuboid, 0),
	}
}

func newCuboidDefault() *cuboid {
	return &cuboid{negatives: make([]*cuboid, 0)}
}

func (c *cuboid) isIntersecting(otherC *cuboid) bool {
	return isIntersecting(c.xMin, c.xMax, otherC.xMin, otherC.xMax) && isIntersecting(c.yMin, c.yMax, otherC.yMin, otherC.yMax) && isIntersecting(c.zMin, c.zMax, otherC.zMin, otherC.zMax)
}

func (c *cuboid) storeNegatives(otherC *cuboid) {
	if c.isIntersecting(otherC) {
		// Double negative
		for _, negative := range c.negatives {
			negative.storeNegatives(otherC)
		}

		iXMin, iXMax := getIntersection(c.xMin, c.xMax, otherC.xMin, otherC.xMax)
		iYMin, iYMax := getIntersection(c.yMin, c.yMax, otherC.yMin, otherC.yMax)
		iZMin, iZMax := getIntersection(c.zMin, c.zMax, otherC.zMin, otherC.zMax)

		c.negatives = append(c.negatives, newCuboid(iXMin, iXMax, iYMin, iYMax, iZMin, iZMax))
	}
}

func (c *cuboid) getVolume() int64 {
	volume := int64(c.xMax-c.xMin+1) * int64(c.yMax-c.yMin+1) * int64(c.zMax-c.zMin+1)

	for _, negative := range c.negatives {
		volume -= negative.getVolume()
	}

	return volume
}

func isIntersecting(x0Min, x0Max, x1Min, x1Max int) bool {
	// Note: if you use < instead of <=, we miss line-overlaps which sometimes can occur
	return (x0Min <= x1Min && x1Min <= x0Max) || (x0Min <= x1Max && x1Max <= x0Max) || (x1Min <= x0Min && x0Min <= x1Max) || (x1Min <= x0Max && x0Max <= x1Max)
}

func getIntersection(x0Min, x0Max, x1Min, x1Max int) (int, int) {
	iMin := x0Min
	if x0Min < x1Min {
		iMin = x1Min
	}

	iMax := x1Max
	if x0Max < x1Max {
		iMax = x0Max
	}

	return iMin, iMax
}

func parseCuboid(lines []string) []*cuboid {
	cuboids := make([]*cuboid, len(lines))
	var err error

	for i, line := range lines {
		cuboids[i] = newCuboidDefault()

		if line[:2] == "on" {
			cuboids[i].isOn = true
			line = strings.Replace(line, "on ", "", 1)
		} else {
			cuboids[i].isOn = false
			line = strings.Replace(line, "off ", "", 1)
		}

		coordinates := strings.Split(line, ",")

		coordinates[0] = strings.Replace(coordinates[0], "x=", "", 1)
		xMinMax := strings.Split(coordinates[0], "..")
		cuboids[i].xMin, err = strconv.Atoi(xMinMax[0])
		panic.Check(err)
		cuboids[i].xMax, err = strconv.Atoi(xMinMax[1])
		panic.Check(err)

		coordinates[1] = strings.Replace(coordinates[1], "y=", "", 1)
		yMinMax := strings.Split(coordinates[1], "..")
		cuboids[i].yMin, err = strconv.Atoi(yMinMax[0])
		panic.Check(err)
		cuboids[i].yMax, err = strconv.Atoi(yMinMax[1])
		panic.Check(err)

		coordinates[2] = strings.Replace(coordinates[2], "z=", "", 1)
		zMinMax := strings.Split(coordinates[2], "..")
		cuboids[i].zMin, err = strconv.Atoi(zMinMax[0])
		panic.Check(err)
		cuboids[i].zMax, err = strconv.Atoi(zMinMax[1])
		panic.Check(err)
	}

	return cuboids
}
