package main

import "fmt"

func main() {
	var (
		blue = [3]int{6, 9, 3}
		red  = [3]int{6, 9, 3}

		// blue = [...]int{6, 9, 3}
		// red  = [3]int{6, 9, 3}

		// blue = [3]int{6, 9, 3}
		// red  = [3]int{3, 9, 6}

		// blue = [3]int{6, 9}
		// red  = [3]int{6, 9, 3}

		// not comparable (type mismatch: length):
		// blue = [3]int{6, 9, 3}
		// red  = [5]int{6, 9, 3}

		// not comparable (type mismatch: element type):
		// blue = [3]int64{6, 9, 3}
		// red  = [3]int{6, 9, 3}
	)
	fmt.Printf("blue bookcase : %v\n", blue)
	fmt.Printf("red bookcase  : %v\n", red)

	fmt.Println("Are they equal?", blue == red)
}
