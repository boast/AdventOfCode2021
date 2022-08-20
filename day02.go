package main

import (
	"mathutil"
	"os"
	"panic"
	"readers"
	"strconv"
	"strings"
)

func Day02Part1() int {
	file, err := os.Open("assets/day02.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	p := mathutil.Point{X: 0, Y: 0}

	for _, s := range lines {
		instructions := strings.Split(s, " ")
		direction := instructions[0]
		length, err := strconv.Atoi(instructions[1])
		panic.Check(err)

		switch direction {
		case "forward":
			p.X += length
			break
		case "down":
			p.Y += length
			break
		case "up":
			p.Y -= length
		}
	}

	return p.X * p.Y
}

func Day02Part2() int {
	file, err := os.Open("assets/day02.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	p := mathutil.Point{X: 0, Y: 0}
	aim := 0

	for _, s := range lines {
		instructions := strings.Split(s, " ")
		direction := instructions[0]
		length, err := strconv.Atoi(instructions[1])
		panic.Check(err)

		switch direction {
		case "forward":
			p.X += length
			p.Y += aim * length
			break
		case "down":
			aim += length
			break
		case "up":
			aim -= length
		}
	}

	return p.X * p.Y
}
