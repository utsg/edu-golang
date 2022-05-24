package coursera

import (
	"fmt"
	"strings"
)

func parseString(userString string) bool {
	userString = strings.ToLower(userString)

	if userString[0:1] == "i" && userString[len(userString)-1:] == "n" && strings.Contains(userString, "a") {
		return true
	}
	return false
}

func main() {
	var userString string

	fmt.Println("Enter a string:")

	_, err := fmt.Scanf("%s", &userString)
	if err != nil {
		fmt.Println(err)
	} else {
		if parseString(userString) {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}
