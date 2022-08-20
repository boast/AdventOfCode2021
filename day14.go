package main

import (
	"math"
	"os"
	"panic"
	"readers"
	"strings"
)

func Day14Part1() int64 {
	file, err := os.Open("assets/day14.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	template := lines[0]
	pairInsertionRules := GetPairInsertionRules(lines)

	occurrences := InitOccurrences(template)

	for i := 0; i < 10; i++ {
		occurrences = Polymerize(occurrences, pairInsertionRules)
	}

	bytesCount := GetBytesCount(occurrences, template[len(template)-1])

	min, max := GetMinMax(bytesCount)

	return max - min
}

func Day14Part2() int64 {
	file, err := os.Open("assets/day14.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	template := lines[0]
	pairInsertionRules := GetPairInsertionRules(lines)

	occurrences := InitOccurrences(template)

	for i := 0; i < 40; i++ {
		occurrences = Polymerize(occurrences, pairInsertionRules)
	}

	bytesCount := GetBytesCount(occurrences, template[len(template)-1])

	min, max := GetMinMax(bytesCount)

	return max - min
}

func InitOccurrences(template string) map[string]int64 {
	occurrences := make(map[string]int64)

	for i := 0; i < len(template)-1; i++ {
		pair := string(template[i]) + string(template[i+1])
		occurrences[pair]++
	}

	return occurrences
}

func Polymerize(occurrences map[string]int64, pairInsertionRules map[string]string) map[string]int64 {
	occurrencesNext := make(map[string]int64)

	// The mapping could be cached as it is constant, but the lookups are really fast as-is
	for pair, count := range occurrences {
		rule := pairInsertionRules[pair]
		occurrencesNext[string(pair[0])+rule] += count
		occurrencesNext[rule+string(pair[1])] += count
	}

	return occurrencesNext
}

func GetBytesCount(occurrences map[string]int64, last byte) map[byte]int64 {
	characterCount := make(map[byte]int64)

	for pair, count := range occurrences {
		characterCount[pair[0]] += count
	}

	// There are only four hard problems in IT / Computational Sciences:
	// 1. Naming things
	// 2. Cache invalidation
	// 3. Off-by-one errors
	characterCount[last]++

	return characterCount
}

func GetMinMax(byteMap map[byte]int64) (int64, int64) {
	var min int64 = math.MaxInt64
	var max int64 = math.MinInt64

	for b := range byteMap {
		if byteMap[b] < min {
			min = byteMap[b]
		}
		if byteMap[b] > max {
			max = byteMap[b]
		}
	}

	return min, max
}

func GetPairInsertionRules(lines []string) map[string]string {
	pairInsertionRules := make(map[string]string)

	for i := 2; i < len(lines); i++ {
		ruleParts := strings.Split(lines[i], " -> ")
		pairInsertionRules[ruleParts[0]] = ruleParts[1]
	}

	return pairInsertionRules
}
