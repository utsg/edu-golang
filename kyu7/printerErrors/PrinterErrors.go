package kyu7

import "strconv"

func PrinterError(s string) string {
	count := 0
	for _, ch := range s {
		if ch < 97 || ch > 109 {
			count++
		}
	}

	return strconv.Itoa(count) + "/" + strconv.Itoa(len(s))
}
