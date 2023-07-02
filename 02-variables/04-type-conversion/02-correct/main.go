package main

import "fmt"

// order of conversion matters...

func main() {
	speed := 100
	force := 2.5

	fmt.Printf("speed     : %d\n", speed)
	fmt.Printf("conversion: %f\n", float64(speed))
	fmt.Printf("expression: %f\n", float64(speed)*force)

	// TYPE MISMATCH:
	//   speed is int
	//   expression is float64
	// speed = float64(speed) * force

	// CORRECT: int * int
	speed = int(float64(speed) * force)

	fmt.Println(speed)
}
