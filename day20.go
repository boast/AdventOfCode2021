package main

import (
	"math"
	"mathutil"
	"os"
	"panic"
	"readers"
)

type bounds struct {
	xMin, xMax int
	yMin, yMax int
}

func newBounds() bounds {
	return bounds{
		xMin: math.MaxInt,
		xMax: math.MinInt,
		yMin: math.MaxInt,
		yMax: math.MinInt,
	}
}

func Day20Part1() int {
	file, err := os.Open("assets/day20.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	lookup := lines[0]
	display := make(map[mathutil.Point]struct{})

	for i := 2; i < len(lines); i++ {
		y := i - 2
		for x := 0; x < len(lines[i]); x++ {
			if string(lines[i][x]) == "#" {
				display[mathutil.Point{X: x, Y: y}] = struct{}{}
			}
		}
	}

	for i := 0; i < 2; i++ {
		display = enhance(display, lookup, i%2 == 0)
	}

	return len(display)
}

func Day20Part2() int {
	file, err := os.Open("assets/day20.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	lookup := lines[0]
	display := make(map[mathutil.Point]struct{})

	for i := 2; i < len(lines); i++ {
		y := i - 2
		for x := 0; x < len(lines[i]); x++ {
			if string(lines[i][x]) == "#" {
				display[mathutil.Point{X: x, Y: y}] = struct{}{}
			}
		}
	}

	for i := 0; i < 50; i++ {
		display = enhance(display, lookup, i%2 == 0)
	}

	return len(display)
}

func enhance(display map[mathutil.Point]struct{}, lookup string, on bool) map[mathutil.Point]struct{} {
	newDisplay := make(map[mathutil.Point]struct{}, len(display))
	b := getBounds(display)

	for y := b.yMin - 2; y <= b.yMax+2; y++ {
		for x := b.xMin - 2; x <= b.xMax+2; x++ {
			pixel := mathutil.Point{X: x, Y: y}
			adjacent := pixel.GetAdjacentGrid()
			binary := 0
			bit := float64(8)

			for i := 0; i < len(adjacent); i++ {
				if _, exists := display[adjacent[i]]; exists == on {
					binary += int(math.Pow(2, bit))
				}
				bit--
			}

			// boolean XOR trick instead of
			// (lookup[binary] == '#' && !on) || (lookup[binary] == '.' && on)
			if (lookup[binary] == '#') != on {
				newDisplay[pixel] = struct{}{}
			}
		}
	}

	return newDisplay
}

func getBounds(display map[mathutil.Point]struct{}) bounds {
	b := newBounds()

	for p := range display {
		if p.X > b.xMax {
			b.xMax = p.X
		}
		if p.X < b.xMin {
			b.xMin = p.X
		}
		if p.Y > b.yMax {
			b.yMax = p.Y
		}
		if p.Y < b.yMin {
			b.yMin = p.Y
		}
	}

	return b
}
