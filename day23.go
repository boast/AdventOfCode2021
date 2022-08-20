package main

import (
	"math"
	"os"
	"panic"
	"readers"
	"strconv"
)

var costs = []int64{1, 10, 100, 1000}

func getCost(unit, from, to int) int64 {
	if from > to {
		tmp := from
		from = to
		to = tmp
	}

	depth := int64((to - 3) / 4)
	targetHall := ((to+1)%4)*2 + 3

	discount := int64(0)

	if from == 0 || from == 6 {
		discount++
	}

	distance := int64(math.Abs(float64(2*from-targetHall))) + depth - discount

	return costs[getType(unit)] * distance
}

type burrowState struct {
	cost      int64
	positions []int
}

func (b *burrowState) moveUnit(unit, position int, newCost int64) burrowState {
	newPositions := make([]int, len(b.positions))
	copy(newPositions, b.positions)
	newPositions[unit] = position

	return burrowState{
		positions: newPositions,
		cost:      b.cost + newCost,
	}
}

func (b *burrowState) isDone() bool {
	for i := 0; i < len(b.positions); i++ {
		if b.positions[i] < 7 || (b.positions[i]+1)%4 != getType(i) {
			return false
		}
	}

	return true
}

func (b *burrowState) toStateString() string {
	grid := make([]int, totalAmphipods()+7)

	for i := 0; i < len(grid); i++ {
		grid[i] = -1
	}
	for i := 0; i < totalAmphipods(); i++ {
		grid[b.positions[i]] = i
	}
	output := ""

	for i := 0; i < len(grid); i++ {
		unitType := getType(grid[i])
		if unitType == -1 {
			output += "-"
		} else {
			output += strconv.Itoa(unitType)
		}
	}

	return output
}

var columnSize = 0

func totalAmphipods() int {
	return columnSize * 4
}

func getType(unit int) int {
	if unit == -1 {
		return -1
	}

	return unit / columnSize
}

func findValidPosition(positions []int, unit int) []bool {
	if positions[unit] < 7 {
		return findValidRoomPositions(positions, unit)
	}
	return findValidHallwayPositions(positions, unit)
}

func findValidHallwayPositions(position []int, unit int) []bool {
	grid := make([]int, totalAmphipods()+7)

	for i := 0; i < len(grid); i++ {
		grid[i] = -1
	}
	for i := 0; i < totalAmphipods(); i++ {
		grid[position[i]] = i
	}

	validPositions := make([]bool, 7)

	currentPosition := position[unit]

	for i := currentPosition - 4; i > 6; i -= 4 {
		if grid[i] != -1 {
			return validPositions
		}
	}

	currentType := getType(unit)

	if (currentPosition+1)%4 == currentType {
		hasToMove := false
		for i := currentPosition + 4; i < totalAmphipods()+7; i += 4 {
			if getType(grid[i]) != currentType {
				hasToMove = true
				break
			}
		}
		if !hasToMove {
			return validPositions
		}
	}

	newPosition := currentPosition

	for newPosition > 10 {
		newPosition -= 4
	}

	for i := 0; i < 7; i++ {
		validPositions[i] = grid[i] == -1 && checkHallwayToRoom(i, newPosition, grid)
	}

	return validPositions
}

func findValidRoomPositions(position []int, unit int) []bool {
	grid := make([]int, totalAmphipods()+7)

	for i := 0; i < len(grid); i++ {
		grid[i] = -1
	}
	for i := 0; i < totalAmphipods(); i++ {
		grid[position[i]] = i
	}

	validPositions := make([]bool, totalAmphipods()+7)

	currentPosition := position[unit]
	currentType := getType(unit)
	roomForType := currentType + 7

	if !checkHallwayToRoom(currentPosition, roomForType, grid) {
		return validPositions
	}
	target := roomForType

	for i := 0; i < columnSize; i++ {
		if grid[roomForType+4*i] == -1 {
			target = roomForType + 4*i
		} else if getType(grid[roomForType+4*i]) != currentType {
			return validPositions
		}
	}

	validPositions[target] = true

	return validPositions
}

func checkHallwayToRoom(hallwayPosition int, roomPosition int, grid []int) bool {
	min := int(math.Min(float64(hallwayPosition+1), float64(roomPosition-5)))
	max := int(math.Max(float64(hallwayPosition-1), float64(roomPosition-6)))

	for i := min; i <= max; i++ {
		if grid[i] != -1 {
			return false
		}
	}
	return true
}

func solve(lines []string) int64 {
	columnSize = len(lines) - 3

	positions := make([]int, totalAmphipods())

	for i := 0; i < columnSize; i++ {
		line := lines[i+2]
		for j := 0; j < 4; j++ {
			unit := int(line[2*j+3]-'A') * columnSize
			for positions[unit] != 0 {
				unit++
			}
			positions[unit] = 4*i + j + 7
		}
	}

	stateQueue := make([]burrowState, 1)
	stateQueue[0] = burrowState{positions: positions, cost: 0}
	cache := make(map[string]int64)

	bestCost := int64(math.MaxInt64)

	for len(stateQueue) > 0 {
		state := stateQueue[0]
		stateQueue = stateQueue[1:]

		if state.cost >= bestCost {
			continue // break if priority queue, as only more expensive can come
		}

		for unit := 0; unit < totalAmphipods(); unit++ {
			for i, validPosition := range findValidPosition(state.positions, unit) {
				if !validPosition {
					continue
				}

				newCost := getCost(unit, state.positions[unit], i)
				newState := state.moveUnit(unit, i, newCost)

				if newState.isDone() {
					if newState.cost < bestCost {
						bestCost = newState.cost
					}
				} else {
					cacheKey := newState.toStateString()
					if cacheValue, exists := cache[cacheKey]; !exists || newState.cost < cacheValue {
						cache[cacheKey] = newState.cost
						stateQueue = append(stateQueue, newState)
					}
				}
			}
		}
	}

	return bestCost
}

func Day23Part1() int64 {
	file, err := os.Open("assets/day23.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	return solve(lines)
}

func Day23Part2() int64 {
	file, err := os.Open("assets/day23part2.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	return solve(lines)
}
