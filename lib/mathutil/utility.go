package mathutil

func Factorial(n int) uint64 {
	factorial := uint64(1)

	for i := 1; i <= n; i++ {
		factorial *= uint64(i)
	}

	return factorial
}

func SumUpTo(n int) int {
	return (n * (n + 1)) / 2
}
