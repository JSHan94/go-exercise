package main

import (
	"fmt"
	"math"
)

func main() {
	radius := 2.0
	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("Area of circle: %.2f\n", area)

	exp := math.Exp(1)
	fmt.Printf("e^1: %.2f\n", exp)

	sqrt := math.Sqrt(10)
	fmt.Printf("Square root of 10: %.2f\n", sqrt)

	abs := math.Abs(-10)
	fmt.Printf("Absolute value of -10: %.2f\n", abs)
}
