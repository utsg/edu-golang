package kyu8

import "fmt"

func SumCubes(n int) int {
	var sum int = 0
	for i := 1; i <= n; i++ {
		sum += i * i * i
	}
	return sum
}

func printTest() {
	fmt.Println(SumCubes(2))
	fmt.Println(SumCubes(3))
}
