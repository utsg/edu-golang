package coursera

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(a int, b int) (int, int) {
	return b, a
}

func BubbleSort(array *[]int) {
	for i := 0; i < len(*array)-1; i++ {
		for j := i + 1; j < len(*array); j++ {
			if (*array)[j] < (*array)[i] {
				(*array)[i], (*array)[j] = Swap((*array)[i], (*array)[j])
			}
		}
	}
}

func InputArray() []int {
	input := bufio.NewReader(os.Stdin)
	bytes, _, err := input.ReadLine()
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	stringFromBytes := strings.Split(string(bytes), " ")
	result := make([]int, len(stringFromBytes))
	for ind, el := range stringFromBytes {
		result[ind], err = strconv.Atoi(el)
		if err != nil {
			fmt.Println("Conversion error", err)
			os.Exit(1)
		}
	}
	return result
}

func main() {
	fmt.Print("Input ints splitted by space\nUnsorted: ")
	unsorted := InputArray()

	BubbleSort(&unsorted)
	fmt.Println("Sorted:", unsorted)
}
