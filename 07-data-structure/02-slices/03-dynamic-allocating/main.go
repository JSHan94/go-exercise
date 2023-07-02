package main

import "fmt"

func main() {
	mySlice := make([]int, 5)

	for i := 0; i < len(mySlice); i++ {
		mySlice[i] = i + 1
	}

	fmt.Println(mySlice)

	yourSlice := mySlice
	mySlice[0] = 100
	// mySlice = append(mySlice, 6)
	fmt.Println(mySlice)
	fmt.Println(yourSlice)
}
