package main

import "fmt"

func main() {
	myMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
		"D": "Delta",
		"E": "Echo",
		"F": "Foxtrot",
	}

	// WARNING : map is unordered!!
	for key, val := range myMap {
		fmt.Println(key, val)
		if key == "B" {
			fmt.Println("found Banana!")
			return
		}
	}
}
