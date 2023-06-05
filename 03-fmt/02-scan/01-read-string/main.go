package main

import (
	"fmt"
)

func main() {
	var a int
	var b string

	// fmt.Scan()
	fmt.Print("Enter a number: ")
	fmt.Scan(&a)
	fmt.Println("You entered:", a)

	// fmt.Scanln()
	fmt.Print("Enter a string: ")
	fmt.Scanln(&b)
	fmt.Println("You entered:", b)

	// fmt.Scanf()
	fmt.Print("Enter a number and a string (separated by a space): ")
	fmt.Scanf("%d %s", &a, &b)
	fmt.Printf("You entered number: %d and string: %s\n", a, b)
}
