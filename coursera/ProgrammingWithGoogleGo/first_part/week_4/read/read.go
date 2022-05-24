package coursera

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	firstName  string
	secondName string
}

func main() {
	var filename string
	var objName Name
	var resultSlice = make([]Name, 0)

	fmt.Println("Enter a filename:")
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		if len(s[0]) > 20 {
			s[0] = s[0][0:20]
		}
		if len(s[1]) > 20 {
			s[1] = s[1][0:20]
		}

		objName.firstName, objName.secondName = s[0], s[1]
		resultSlice = append(resultSlice, objName)
	}

	for _, name := range resultSlice {
		fmt.Println(name.firstName, name.secondName)
	}
}
