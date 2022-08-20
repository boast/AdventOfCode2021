package main

import (
	"collections"
	"os"
	"panic"
	"readers"
	"strconv"
	"strings"
)

func Day06Part1() int {
	file, err := os.Open("assets/day06.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	lanternFishStrings := strings.Split(lines[0], ",")
	lanternFishGenerations := ParseLanternFish(lanternFishStrings)

	return GetLanternFishAfterGenerations(lanternFishGenerations, 80)
}

func Day06Part2() int {
	file, err := os.Open("assets/day06.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	lanternFishStrings := strings.Split(lines[0], ",")
	lanternFishGenerations := ParseLanternFish(lanternFishStrings)

	return GetLanternFishAfterGenerations(lanternFishGenerations, 256)
}

func ParseLanternFish(lanternFishStrings []string) []int {
	lanternFishGenerations := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, lanternFishString := range lanternFishStrings {
		i, err := strconv.Atoi(lanternFishString)
		panic.Check(err)
		lanternFishGenerations[i]++
	}

	return lanternFishGenerations
}

func GetLanternFishAfterGenerations(lanternFishGenerations []int, generations int) int {
	for i := 0; i < generations; i++ {
		newGeneration := lanternFishGenerations[7]
		lanternFishGenerations[7] = lanternFishGenerations[8]
		lanternFishGenerations[8] = lanternFishGenerations[i%7]
		lanternFishGenerations[i%7] += newGeneration
	}

	return collections.SumInt(lanternFishGenerations)
}
