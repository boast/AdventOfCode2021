package main

import (
	"os"
	"panic"
	"readers"
	"strconv"
)

func Day03Part1() int64 {
	file, err := os.Open("assets/day03.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	positions := len(lines[0])
	gamma := ""
	epsilon := ""

	for p := 0; p < positions; p++ {
		var count1Bits = 0
		var count0Bits = 0

		for _, line := range lines {
			var b = line[p]

			if b == '1' {
				count1Bits++
			} else {
				count0Bits++
			}
		}

		if count1Bits > count0Bits {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaInt, err := strconv.ParseInt(gamma, 2, 64)
	panic.Check(err)
	epsilonInt, err := strconv.ParseInt(epsilon, 2, 64)
	panic.Check(err)

	return gammaInt * epsilonInt
}

func Day03Part2() int64 {
	file, err := os.Open("assets/day03.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	lines2 := make([]string, len(lines))
	copy(lines2, lines)

	p := 0
	for len(lines) > 1 {
		count0Bits := 0
		count1Bits := 0

		for i := 0; i < len(lines); i++ {
			if lines[i][p] == '1' {
				count1Bits++
			} else {
				count0Bits++
			}
		}

		var toRemove uint8 = '0'

		if count0Bits > count1Bits {
			toRemove = '1'
		}

		for i := 0; i < len(lines); i++ {
			if lines[i][p] == toRemove {
				lines = removeUnordered(lines, i)
				i--
			}
		}

		p++
	}

	oxygen := lines[0]

	p = 0
	for len(lines2) > 1 {
		count0Bits := 0
		count1Bits := 0

		for i := 0; i < len(lines2); i++ {
			if lines2[i][p] == '1' {
				count1Bits++
			} else {
				count0Bits++
			}
		}

		var toRemove uint8 = '1'

		if count1Bits < count0Bits {
			toRemove = '0'
		}

		for i := 0; i < len(lines2); i++ {
			if lines2[i][p] == toRemove {
				lines2 = removeUnordered(lines2, i)
				i--
			}
		}

		p++
	}

	co2 := lines2[0]

	oxygenInt, err := strconv.ParseInt(oxygen, 2, 64)
	panic.Check(err)
	co2Int, err := strconv.ParseInt(co2, 2, 64)
	panic.Check(err)

	return oxygenInt * co2Int
}

func removeUnordered(s []string, i int) []string {
	// Example: Remove element C from slice (order does not matter)
	// A, B, C, D, E
	// A, B, E, D, E -- Copy last (E) to element position
	// A, B, E, D    -- Remove last element
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
