package main

import (
	"os"
	"panic"
	"readers"
)

func Day01Part1() int {
	file, err := os.Open("assets/day01.txt")
	panic.Check(err)

	integers, err := readers.ReadIntegers(file)
	panic.Check(err)

	count := 0

	for i := 1; i < len(integers); i++ {
		var current = integers[i]
		var previous = integers[i-1]

		if current > previous {
			count++
		}
	}

	return count
}

func Day01Part2() int {
	file, err := os.Open("assets/day01.txt")
	panic.Check(err)

	integers, err := readers.ReadIntegers(file)
	panic.Check(err)

	count := 0

	for i := 1; i < len(integers)-2; i++ {
		var current = integers[i] + integers[i+1] + integers[i+2]
		var previous = integers[i-1] + integers[i] + integers[i+1]

		if current > previous {
			count++
		}
	}

	return count
}
