package mathutil

import "sort"

func AverageInt(slice []int) float64 {
	sum := 0

	for _, i := range slice {
		sum += i
	}

	return (float64(sum)) / (float64(len(slice)))
}

func MedianInt(slice []int) float64 {
	sort.Ints(slice)

	count := len(slice)
	middle := count / 2

	if IntIsOdd(count) {
		return float64(slice[middle])
	}

	return float64((slice[middle-1] + slice[middle]) / 2)
}

func IntIsEven(i int) bool {
	return i%2 == 0
}

func IntIsOdd(i int) bool {
	return i%2 != 0
}
