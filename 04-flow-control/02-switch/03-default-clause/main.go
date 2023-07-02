package main

import (
	"fmt"
)

func main() {
	city := "Paris"

	switch city {
	case "Paris":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	default:
		fmt.Println("Where?")
		// there can be only one default in a switch
		// default:
		// 	fmt.Println("Where?")
	}

	// if city == "Paris" {
	// 	fmt.Println("France")
	// } else if city == "Tokyo" {
	// 	fmt.Println("Japan")
	// } else {
	// 	fmt.Println("Where?")
	// }
}
