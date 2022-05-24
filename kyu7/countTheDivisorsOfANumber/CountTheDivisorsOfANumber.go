package kyu7

func Divisors(n int) int {
	count := 1

	for i := 1; i < n; i++ {
		if (n % i) == 0 {
			count++
		}
	}

	return count
}
