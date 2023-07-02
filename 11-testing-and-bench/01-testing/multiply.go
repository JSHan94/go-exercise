package multiply

import "fmt"

func Multiply(a, b int) int {
	if a < 0 {
		fmt.Println("a is bigger than 0")
	}
	return a * b
}
