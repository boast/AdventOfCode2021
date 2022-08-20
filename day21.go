package main

import (
	"os"
	"panic"
	"readers"
	"strconv"
	"strings"
)

func Day21Part1() int {
	file, err := os.Open("assets/day21.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	player1Position, err := strconv.Atoi(strings.Replace(lines[0], "Player 1 starting position: ", "", 1))
	panic.Check(err)
	player2Position, err := strconv.Atoi(strings.Replace(lines[1], "Player 2 starting position: ", "", 1))
	panic.Check(err)

	var player1Points, player2Points, rolls int
	d := die{value: 1, size: 100}

	for true {
		player1Position += d.Roll() + d.Roll() + d.Roll()
		player1Position = (player1Position-1)%10 + 1
		rolls += 3
		player1Points += player1Position

		if player1Points >= 1000 {
			break
		}
		player2Position += d.Roll() + d.Roll() + d.Roll()
		player2Position = (player2Position-1)%10 + 1
		rolls += 3
		player2Points += player2Position

		if player2Points >= 1000 {
			break
		}
	}

	if player1Points < player2Points {
		return player1Points * rolls
	} else {
		return player2Points * rolls
	}
}

func Day21Part2() int64 {
	file, err := os.Open("assets/day21.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	player1Position, err := strconv.Atoi(strings.Replace(lines[0], "Player 1 starting position: ", "", 1))
	panic.Check(err)
	player2Position, err := strconv.Atoi(strings.Replace(lines[1], "Player 2 starting position: ", "", 1))
	panic.Check(err)

	cache := make(map[state]score)

	s := quantumGame(&cache, state{player1Position: player1Position, player2Position: player2Position})

	if s.player1Count > s.player2Count {
		return s.player1Count
	} else {
		return s.player2Count
	}
}

func quantumGame(cache *map[state]score, s state) score {
	if s.player2Score >= 21 {
		return score{player2Count: 1}
	}

	if cachedScore, exists := (*cache)[s]; exists {
		return cachedScore
	}

	currentScore := score{}

	// We roll all possible outcomes at the same time
	// see: https://old.reddit.com/r/adventofcode/comments/rl6p8y/2021_day_21_solutions/hpe4z64/
	for _, distribution := range totalDieDistribution {
		newPosition := (s.player1Position+distribution.dieValue-1)%10 + 1

		// Swap active player
		subScore := quantumGame(cache, state{
			player1Score:    s.player2Score,
			player2Score:    s.player1Score + newPosition,
			player1Position: s.player2Position,
			player2Position: newPosition,
		})

		// Swap new scores as well
		currentScore.player1Count += subScore.player2Count * distribution.times
		currentScore.player2Count += subScore.player1Count * distribution.times
	}

	(*cache)[s] = currentScore

	return currentScore
}

var totalDieDistribution = []dieDistribution{
	{dieValue: 3, times: 1},
	{dieValue: 4, times: 3},
	{dieValue: 5, times: 6},
	{dieValue: 6, times: 7},
	{dieValue: 7, times: 6},
	{dieValue: 8, times: 3},
	{dieValue: 9, times: 1},
}

type dieDistribution struct {
	dieValue int
	times    int64
}

type state struct {
	player1Score    int
	player2Score    int
	player1Position int
	player2Position int
}

type score struct {
	player1Count int64
	player2Count int64
}

type die struct {
	value int
	size  int
}

func (d *die) Roll() int {
	next := d.value

	d.value++

	if d.value > d.size {
		d.value = 1
	}

	return next
}
