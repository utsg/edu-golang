package kyu7

import "fmt"

func findMax(dict map[rune]int) rune {
	max := 0
	var result rune
	for ch, counter := range dict {
		if counter > max {
			max = counter
			result = ch
		} else if counter == max && ch < result {
			result = ch
		}
	}

	return result
}

func Solve(s string) rune {
	dict := make(map[rune]int)
	for _, ch := range s {
		if _, ok := dict[ch]; ok {
			dict[ch]++
		} else {
			dict[ch] = 1
		}
	}

	fmt.Println(s + "[" + string(findMax(dict)) + "]")
	return findMax(dict)
}
