package kyu7

import (
	"strings"
)

func ToJadenCase(str string) string {
	words := strings.Split(str, " ")
	
	for i, word := range words {
		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}

	return strings.Join(words, " ")
}
