package kyu7

import (
	"strings"
)

func reverseWord(str string) (result string) {
	for _, ch := range str {
		result = string(ch) + result
	}
	return
}

func ReverseWords(str string) string {
	words := strings.Split(str, " ")

	for i, v := range words{
	  words[i] = reverseWord(string(v))
	}
  
	return strings.Join(words, " ") // reverse those words
}
