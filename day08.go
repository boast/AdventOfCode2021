package main

import (
	"collections"
	"math"
	"os"
	"panic"
	"readers"
	"strings"
)

func Day08Part1() int {
	file, err := os.Open("assets/day08.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	sum := 0

	for _, line := range lines {
		parts := strings.Split(line, " | ")
		outputValueStrings := strings.Split(parts[1], " ")

		for _, valueString := range outputValueStrings {
			switch len(valueString) {
			case 2: // 1
				fallthrough
			case 3: // 7
				fallthrough
			case 4: // 4
				fallthrough
			case 7: // 8
				sum++
				break
			}
		}
	}

	return sum
}

func Day08Part2() int {
	file, err := os.Open("assets/day08.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	sum := 0

	for _, line := range lines {
		parts := strings.Split(line, " | ")

		signalStrings := strings.Split(parts[0], " ")
		outputValueStrings := strings.Split(parts[1], " ")

		numbers := make([][]byte, 10)

		// Find 1
		for idx, valueString := range signalStrings {
			if len(valueString) == 2 {
				numbers[1] = []byte(valueString)
				signalStrings = collections.RemoveFromString(signalStrings, idx)
				break
			}
		}

		// Find 4
		for idx, valueString := range signalStrings {
			if len(valueString) == 4 {
				numbers[4] = []byte(valueString)
				signalStrings = collections.RemoveFromString(signalStrings, idx)
				break
			}
		}

		// Find 7
		for idx, valueString := range signalStrings {
			if len(valueString) == 3 {
				numbers[7] = []byte(valueString)
				signalStrings = collections.RemoveFromString(signalStrings, idx)
				break
			}
		}

		// Find 8
		for idx, valueString := range signalStrings {
			if len(valueString) == 7 {
				numbers[8] = []byte(valueString)
				signalStrings = collections.RemoveFromString(signalStrings, idx)
				break
			}
		}

		// Find 3
		for idx, valueString := range signalStrings {
			if numbers[3] = []byte(valueString); len(valueString) == 5 && len(collections.DifferenceByte(numbers[3], numbers[1])) == 3 {
				signalStrings = collections.RemoveFromString(signalStrings, idx)
				break
			}
		}

		// Find 2
		for idx, valueString := range signalStrings {
			if numbers[2] = []byte(valueString); len(valueString) == 5 && len(collections.DifferenceByte(numbers[2], numbers[4])) == 5 {
				signalStrings = collections.RemoveFromString(signalStrings, idx)
				break
			}
		}

		// Find 5
		for idx, valueString := range signalStrings {
			if numbers[5] = []byte(valueString); len(valueString) == 5 {
				signalStrings = collections.RemoveFromString(signalStrings, idx)
				break
			}
		}

		// Find 6
		for idx, valueString := range signalStrings {
			if numbers[6] = []byte(valueString); len(collections.DifferenceByte(numbers[6], numbers[1])) == 6 {
				signalStrings = collections.RemoveFromString(signalStrings, idx)
				break
			}
		}

		// Find 9

		for idx, valueString := range signalStrings {
			if numbers[9] = []byte(valueString); len(collections.DifferenceByte(numbers[9], numbers[4])) == 2 {
				signalStrings = collections.RemoveFromString(signalStrings, idx)
				break
			}
		}

		// Find 0
		numbers[0] = []byte(signalStrings[0])

		output := 0
		for idx, valueString := range outputValueStrings {
			valueStringAsBytes := []byte(valueString)

			for i := 0; i < 10; i++ {
				if len(numbers[i]) == len(valueStringAsBytes) && len(collections.DifferenceByte(numbers[i], valueStringAsBytes)) == 0 {
					output += int(math.Pow(10, float64(3-idx))) * i
					continue
				}
			}
		}

		sum += output
	}

	return sum
}
