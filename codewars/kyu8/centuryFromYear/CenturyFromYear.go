package kyu8

func century(year int) int {
	if year < 100 {
		return 1
	} else if year%100 == 0 {
		return year / 100
	}
	return year/100 + 1
}

func printTest() {
	println(century(1990))
	println(century(1705))
	println(century(1900))
	println(century(1601))
	println(century(2000))
	println(century(89))
}
