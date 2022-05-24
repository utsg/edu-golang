package kyu8

import (
	"fmt"
	"strconv"
	"strings"
)

func Factorial(i int) int {
	if i == 0 {
		return 1
	}
	return i * Factorial(i-1)
}

func Strong(n int) string {
	var sum int = 0

	for _, digit := range strings.Split(strconv.Itoa(n), "") {
		digitConverted, err := strconv.Atoi(digit)
		if err != nil {
			fmt.Println("some shit goes wrong")
		}
		sum += Factorial(digitConverted)
	}

	if int(sum) == n {
		return "STRONG!!!!"
	}
	return "Not Strong !!"
}

func printTest() {
	fmt.Println(Strong(145))
}
