package main

import (
	"fmt"
)

func main() {
	city := "Paris"

	switch city {
	case "Paris":
		fmt.Println("France")
		// break // unnecessary in Go

		vip := true
		fmt.Println("VIP trip?", vip)
	case "Tokyo":
		fmt.Println("Japan")
		// fmt.Println("VIP trip?", vip)
	}

	// if city == "Paris" {
	// 	fmt.Println("France")
	// } else if city == "Tokyo" {
	// 	fmt.Println("Japan")
	// }
}
