package main

import "fmt"

func main() {
	a := false
	b := true

	// Short Circuit
	if a && b {
		fmt.Println("This won't be printed")
	}

	if num := 10; num%2 == 0 {
		fmt.Println("Even number")
	}
}
