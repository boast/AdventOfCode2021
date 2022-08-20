package collections

import "strconv"

func MapString(slice []string, mapFunction func(i string) string) []string {
	newSlice := make([]string, len(slice))

	for idx, s := range slice {
		newSlice[idx] = mapFunction(s)
	}

	return newSlice
}

func RemoveFromString(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func RemoveUnorderedFromString(slice []string, index int) []string {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func FilterFromString(slice []string, filter func(element string) bool) []string {
	for index, element := range slice {
		if filter(element) {
			slice = RemoveFromString(slice, index)
		}
	}
	return slice
}

func FilterEmptyFromString(slice []string) []string {
	return FilterFromString(slice, func(element string) bool {
		return element == ""
	})
}

func ContainsString(slice []string, value string) bool {
	for _, element := range slice {
		if element == value {
			return true
		}
	}
	return false
}

func MapInt(slice []int, mapFunction func(i int) int) []int {
	newSlice := make([]int, len(slice))

	for idx, i := range slice {
		newSlice[idx] = mapFunction(i)
	}

	return newSlice
}

func RemoveFromInt(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func RemoveUnorderedFromInt(slice []int, index int) []int {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func FilterFromInt(slice []int, filter func(element int) bool) []int {
	for index, element := range slice {
		if filter(element) {
			slice = RemoveFromInt(slice, index)
		}
	}
	return slice
}

func FilterValueFromInt(slice []int, value int) []int {
	return FilterFromInt(slice, func(element int) bool {
		return element == value
	})
}

func StringSliceToIntSlice(slice []string) []int {
	intSlice := make([]int, len(slice))

	for i, s := range slice {
		intSlice[i], _ = strconv.Atoi(s)
	}

	return intSlice
}

func SumInt(slice []int) int {
	sum := 0

	for _, n := range slice {
		sum += n
	}

	return sum
}

func SumInt64(slice []int64) int64 {
	var sum int64 = 0

	for _, n := range slice {
		sum += n
	}

	return sum
}

func DifferenceByte(sliceA, sliceB []byte) []byte {
	var difference []byte

	for i := 0; i < 2; i++ {
		for _, elementSliceA := range sliceA {
			found := false

			for _, elementSliceB := range sliceB {
				if elementSliceA == elementSliceB {
					found = true
					break
				}
			}

			if !found {
				difference = append(difference, elementSliceA)
			}
		}

		if i == 0 {
			sliceA, sliceB = sliceB, sliceA
		}
	}

	return difference
}

func ContainsByte(slice []byte, value byte) bool {
	for _, element := range slice {
		if element == value {
			return true
		}
	}
	return false
}
