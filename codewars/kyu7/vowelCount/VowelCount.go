package kyu7

import "strings"

func GetCount(str string) (count int) {
	count = 0
	dict := map[string]int{
		"A": 1,
		"E": 2,
		"I": 3,
		"O": 4,
		"U": 5,
	}

	for _, ch := range str {
		if _, ok := dict[strings.ToUpper(string(ch))]; ok {
			count++
		}
	}

	return count
}

func BestWay(str string) (count int) {
	for _, c := range str {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
