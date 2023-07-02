package main

import "fmt"

func main() {
	score, valid := 4, false

	if score > 3 && valid {
		fmt.Println("true!")
	} else if score > 6 {
		fmt.Println("score is more than 6")
	} else if score > 5 {
		fmt.Println("score is more than 5")
	} else {
		fmt.Println("nothing matched!!")
	}
}
