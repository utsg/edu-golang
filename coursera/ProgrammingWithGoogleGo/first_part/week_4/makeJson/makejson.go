package coursera

import (
	"encoding/json"
	"fmt"
)

//Write a program which prompts the user to first enter a name, and then enter an address.
//Your program should create a map and add the name and address to the map using the keys “name” and “address”,
//respectively.Your program should use Marshal() to create a JSON object from the map, and then your program should
//print the JSON object.

func main() {
	var firstName string
	var address string

	fmt.Println("Enter the first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter the address")
	fmt.Scan(&address)

	resultMap := map[string]string{
		firstName: address,
	}

	jsonString, err := json.Marshal(resultMap)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(string(jsonString))
}
