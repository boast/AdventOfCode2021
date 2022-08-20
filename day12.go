package main

import (
	"collections"
	"os"
	"panic"
	"readers"
	"strings"
)

func Day12Part1() int {
	file, err := os.Open("assets/day12.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	caves := make(map[string][]string)

	for _, line := range lines {
		caveParts := strings.Split(line, "-")

		caves[caveParts[0]] = append(caves[caveParts[0]], caveParts[1])
		caves[caveParts[1]] = append(caves[caveParts[1]], caveParts[0])
	}

	return FindPaths("start", []string{}, caves, false)
}

func Day12Part2() int {
	file, err := os.Open("assets/day12.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	caves := make(map[string][]string)

	for _, line := range lines {
		caveParts := strings.Split(line, "-")

		caves[caveParts[0]] = append(caves[caveParts[0]], caveParts[1])
		caves[caveParts[1]] = append(caves[caveParts[1]], caveParts[0])
	}

	return FindPaths("start", []string{}, caves, true)
}

func FindPaths(current string, visited []string, caves map[string][]string, extra bool) int {
	total := 0

	for _, direction := range caves[current] {
		if direction == "end" {
			total += 1
		} else if direction == strings.ToUpper(direction) {
			total += FindPaths(direction, visited, caves, extra)
		} else if direction != "start" {
			if !collections.ContainsString(visited, direction) {
				total += FindPaths(direction, append(visited, direction), caves, extra)
			} else if extra {
				total += FindPaths(direction, visited, caves, false)
			}
		}
	}

	return total
}
