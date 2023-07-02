package main

import "fmt"

// scope shadowing

func main() {
	x := 5
	fmt.Println(x) // Output: 5
	{
		fmt.Println(x) // Output: 5
		x := 10
		fmt.Println(x) // Output: 10
	}
	fmt.Println(x) // Output: 5
}
