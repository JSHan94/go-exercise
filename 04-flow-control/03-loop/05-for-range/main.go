package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 5, 6}

	for i, num := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}

	// myMap := map[string]string{
	// 	"Apple":  "Red",
	// 	"Banana": "Yellow",
	// 	"Grape":  "Purple",
	// }

	// for key, value := range myMap {
	// 	fmt.Println(key, value)
	// }
}
