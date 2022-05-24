package coursera

import (
	"fmt"
	"sort"
	"strconv"
)

// Write a program which prompts the user to enter integers
// and stores the integers in a sorted slice.
// The program should be written as a loop.
// Before entering the loop, the program should create an empty integer slice of size (length) 3.
// During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
// The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
// The slice must grow in size to accommodate any number of integers which the user decides to enter.
// The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

func main() {
	var res = make([]int, 3)
	fmt.Println("Enter the integers")

	var num string
	fmt.Scan(&num)

	for num != "X" {
		converted, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("something goes wrong")
			fmt.Println(err)
		}
		res = append(res, converted)
		sort.Ints(res)
		fmt.Println(res)
		fmt.Scan(&num)
	}
}
