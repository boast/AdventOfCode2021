package mathutil

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

type PointMap map[Point]interface{}

func (p *Point) Add(q Point) {
	p.X += q.X
	p.Y += q.Y
}

func (p *Point) AddScalar(i int) {
	p.X += i
	p.Y += i
}

func (p *Point) MultiplyScalar(i int) {
	p.X *= i
	p.Y *= i
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func (p *Point) GetNeighbours() []Point {
	return []Point{
		{X: p.X, Y: p.Y + 1}, // N
		{X: p.X + 1, Y: p.Y}, // E
		{X: p.X, Y: p.Y - 1}, // S
		{X: p.X - 1, Y: p.Y}, // W
	}
}

func (p *Point) GetAdjacent() []Point {
	return []Point{
		{X: p.X, Y: p.Y + 1},     // N
		{X: p.X + 1, Y: p.Y + 1}, // NE
		{X: p.X + 1, Y: p.Y},     // E
		{X: p.X + 1, Y: p.Y - 1}, // SE
		{X: p.X, Y: p.Y - 1},     // S
		{X: p.X - 1, Y: p.Y - 1}, // SW
		{X: p.X - 1, Y: p.Y},     // W
		{X: p.X - 1, Y: p.Y + 1}, // NW
	}
}

func (p *Point) GetAdjacentGrid() []Point {
	return []Point{
		{X: p.X - 1, Y: p.Y - 1}, // Top-Left
		{X: p.X, Y: p.Y - 1},     // Top
		{X: p.X + 1, Y: p.Y - 1}, // Top-Right
		{X: p.X - 1, Y: p.Y},     // Center-Left
		{X: p.X, Y: p.Y},         // Center
		{X: p.X + 1, Y: p.Y},     // Center-Right
		{X: p.X - 1, Y: p.Y + 1}, // Bottom-Left
		{X: p.X, Y: p.Y + 1},     // Bottom
		{X: p.X + 1, Y: p.Y + 1}, // Bottom-Right
	}
}

func GetIntMap(lines []string) map[Point]int {
	var err error
	intMap := make(map[Point]int)

	for y, line := range lines {
		for x, value := range strings.Split(line, "") {
			intMap[Point{X: x, Y: y}], err = strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
		}
	}

	return intMap
}

func (m *PointMap) GetMaxCoordinates() (int, int) {
	xMax := 0
	yMax := 0
	for point := range *m {
		if point.X > xMax {
			xMax = point.X
		}
		if point.Y > yMax {
			yMax = point.Y
		}
	}

	return xMax, yMax
}

func (m *PointMap) Print() {
	xMax, yMax := m.GetMaxCoordinates()

	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			if _, exists := (*m)[Point{X: x, Y: y}]; exists {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func (m *PointMap) FoldX(xFold int) {
	xMax, yMax := m.GetMaxCoordinates()

	for x := xFold; x <= xMax; x++ {
		for y := 0; y <= yMax; y++ {
			if _, exists := (*m)[Point{X: x, Y: y}]; exists {
				(*m)[Point{X: xFold - (x - xFold), Y: y}] = struct{}{}
				delete(*m, Point{X: x, Y: y})
			}
		}
	}
}

func (m *PointMap) FoldY(yFold int) {
	xMax, yMax := m.GetMaxCoordinates()

	for y := yFold; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			if _, exists := (*m)[Point{X: x, Y: y}]; exists {
				(*m)[Point{X: x, Y: yFold - (y - yFold)}] = struct{}{}
				delete(*m, Point{X: x, Y: y})
			}
		}
	}
}
