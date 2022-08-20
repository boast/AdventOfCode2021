package main

import (
	"collections"
	"os"
	"panic"
	"readers"
	"sort"
	"strings"
)

func Day10Part1() int {
	file, err := os.Open("assets/day10.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	points := 0

	for _, line := range lines {
		corrupted, _ := GetCorrupted(line)

		switch corrupted {
		case ")":
			points += 3
			break
		case "]":
			points += 57
			break
		case "}":
			points += 1197
			break
		case ">":
			points += 25137
			break
		}
	}

	return points
}

func Day10Part2() int {
	file, err := os.Open("assets/day10.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	var scores []int

	for _, line := range lines {
		corrupted, incomplete := GetCorrupted(line)

		if corrupted == "" {
			score := 0

			for !incomplete.IsEmpty() {
				c, _ := incomplete.Pop()

				switch c {
				case "(":
					score *= 5
					score += 1
					break
				case "[":
					score *= 5
					score += 2
					break
				case "{":
					score *= 5
					score += 3
					break
				case "<":
					score *= 5
					score += 4
					break
				}
			}

			scores = append(scores, score)
		}
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func GetCorrupted(line string) (string, collections.StackString) {
	var stack collections.StackString

	for _, c := range strings.Split(line, "") {
		switch c {
		case "(":
			fallthrough
		case "[":
			fallthrough
		case "{":
			fallthrough
		case "<":
			stack.Push(c)
			break
		case ")":
			if c2, _ := stack.Pop(); c2 != "(" {
				return c, nil
			}
			break
		case "]":
			if c2, _ := stack.Pop(); c2 != "[" {
				return c, nil
			}
			break
		case "}":
			if c2, _ := stack.Pop(); c2 != "{" {
				return c, nil
			}
			break
		case ">":
			if c2, _ := stack.Pop(); c2 != "<" {
				return c, nil
			}
			break
		}
	}

	return "", stack
}
