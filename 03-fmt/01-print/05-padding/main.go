package main

import (
	"fmt"
)

func main() {
	// Right-aligned padding
	fmt.Printf("|%4d|\n", 5)     // Output: |   5|
	fmt.Printf("|%4d|\n", 12345) // Output: |12345|

	// Left-aligned padding
	fmt.Printf("|%-4s|\n", "Go")     // Output: |Go  |
	fmt.Printf("|%-4s|\n", "Golang") // Output: |Golang|

	// Zero padding
	fmt.Printf("|%04d|\n", 5)     // Output: |0005|
	fmt.Printf("|%04d|\n", 12345) // Output: |12345|
}
