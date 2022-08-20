package main

import (
	"os"
	"panic"
	"readers"
)

const (
	cucumberRight = uint8('>')
	cucumberDown  = uint8('v')
	empty         = uint8('.')
)

type cucumberMove struct {
	x, y, next int
}

func Day25Part1() int {
	file, err := os.Open("assets/day25.txt")
	panic.Check(err)

	seafloor, err := readers.ReadStrings(file)
	panic.Check(err)

	steps := 0

	yMax := len(seafloor)
	xMax := len(seafloor[0])

	for true {
		steps++

		didMove := false
		moves := make([]cucumberMove, 0)

		for y := 0; y < yMax; y++ {
			for x := 0; x < xMax; x++ {
				if seafloor[y][x] == cucumberRight {
					next := (x + 1) % xMax
					if seafloor[y][next] == empty {
						moves = append(moves, cucumberMove{x, y, next})
						didMove = true
					}
				}
			}
		}

		for _, move := range moves {
			seafloor[move.y] = replaceAtIndex(seafloor[move.y], empty, move.x)
			seafloor[move.y] = replaceAtIndex(seafloor[move.y], cucumberRight, move.next)
		}

		moves = make([]cucumberMove, 0)

		for x := 0; x < xMax; x++ {
			for y := 0; y < yMax; y++ {
				if seafloor[y][x] == cucumberDown {
					next := (y + 1) % yMax
					if seafloor[next][x] == empty {
						moves = append(moves, cucumberMove{x, y, next})
						didMove = true
					}
				}
			}
		}

		for _, move := range moves {
			seafloor[move.y] = replaceAtIndex(seafloor[move.y], empty, move.x)
			seafloor[move.next] = replaceAtIndex(seafloor[move.next], cucumberDown, move.x)
		}

		if !didMove {
			return steps
		}
	}

	return -1
}

//goland:noinspection GoUnusedFunction
func printSeafloor(seafloor []string) string {
	output := ""

	for _, s := range seafloor {
		output += s + "\n"
	}

	return output
}

func replaceAtIndex(str string, replacement uint8, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func Day25Part2() string {
	return "Marry X-Mas"
}
