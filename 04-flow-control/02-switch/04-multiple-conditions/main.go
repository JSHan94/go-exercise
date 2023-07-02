package main

import (
	"fmt"
)

func main() {
	city := "Lyon"

	switch city {
	case "Paris", "Lyon":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	}

	// if city == "Paris" || city == "Lyon" {
	// 	fmt.Println("France")
	// } else if city == "Tokyo" {
	// 	fmt.Println("Japan")
	// }
}
