package main

import (
	"collections"
	"container/heap"
	"mathutil"
	"os"
	"panic"
	"readers"
)

func Day15Part1() int {
	file, err := os.Open("assets/day15.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	start := mathutil.Point{}
	end := mathutil.Point{X: len(lines[0]) - 1, Y: len(lines) - 1}
	riskMap := mathutil.GetIntMap(lines)

	path := GetPath(riskMap, start, end)

	return GetPathRisk(riskMap, path, start, end)
}

func Day15Part2() int {
	file, err := os.Open("assets/day15.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	xMax := len(lines[0])
	yMax := len(lines)

	start := mathutil.Point{}
	end := mathutil.Point{X: (xMax * 5) - 1, Y: (yMax * 5) - 1}
	riskMap := mathutil.GetIntMap(lines)

	for yMod := 0; yMod < 5; yMod++ {
		for xMod := 0; xMod < 5; xMod++ {
			if yMod == 0 && xMod == 0 {
				continue
			}

			for y := 0; y < yMax; y++ {
				for x := 0; x < xMax; x++ {
					risk := riskMap[mathutil.Point{X: x, Y: y}] + xMod + yMod
					if risk > 9 {
						risk -= 9
					}
					newX := x + (xMod * xMax)
					riskMap[mathutil.Point{X: newX, Y: y + (yMod * yMax)}] = risk
				}
			}
		}
	}

	path := GetPath(riskMap, start, end)

	return GetPathRisk(riskMap, path, start, end)
}

func GetPath(riskMap map[mathutil.Point]int, start, end mathutil.Point) map[mathutil.Point]mathutil.Point {
	pq := make(collections.PriorityQueue, 1)
	pq[0] = &collections.PriorityQueueItem{Value: start}
	heap.Init(&pq)

	// This prevents us to iterate the pq slice manually to check if an item is already in the pq (tradeoff: memory)
	pqLookup := make(map[mathutil.Point]struct{}, 1)
	pqLookup[start] = struct{}{}

	path := make(map[mathutil.Point]mathutil.Point)

	lowestRisk := make(map[mathutil.Point]int)
	lowestRisk[start] = 0

	isInMap := func(point mathutil.Point) bool {
		return point.X >= 0 && point.X <= end.X && point.Y >= 0 && point.Y <= end.Y
	}

	for pq.Len() > 0 {
		// We Pop twice: once from the heap, once from the map
		current := heap.Pop(&pq).(*collections.PriorityQueueItem)
		currentPoint := current.Value.(mathutil.Point)
		delete(pqLookup, currentPoint)

		if currentPoint == end {
			break
		}

		for _, neighbour := range currentPoint.GetNeighbours() {
			if !isInMap(neighbour) {
				continue
			}

			neighbourLowestRiskCandidate := lowestRisk[currentPoint] + riskMap[neighbour]

			if neighbourLowestRisk, existsInLowestRisk := lowestRisk[neighbour]; !existsInLowestRisk || neighbourLowestRiskCandidate < neighbourLowestRisk {
				path[neighbour] = currentPoint
				lowestRisk[neighbour] = neighbourLowestRiskCandidate

				if _, existsInLookup := pqLookup[neighbour]; !existsInLookup {
					// We Push twice: once onto the heap, once onto the map
					heap.Push(&pq, &collections.PriorityQueueItem{Value: neighbour, Priority: lowestRisk[neighbour]})
					pqLookup[neighbour] = struct{}{}
				}
			}
		}

		if pq.Len() == 0 {
			panic.Panic("Could not find solution")
		}
	}

	return path
}

func GetPathRisk(riskMap map[mathutil.Point]int, path map[mathutil.Point]mathutil.Point, start, end mathutil.Point) int {
	risk := 0

	current := end

	for current != start {
		risk += riskMap[current]
		current = path[current]
	}

	return risk
}
