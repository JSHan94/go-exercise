package main

import "fmt"

// Multi return values
func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

// Named return and Naked Return
func divmod2(a int, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return // naked return
}

// Variadic Functions
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Functions as Values and Types
func apply(fn func(int, int) int, a int, b int) int {
	return fn(a, b)
}

func sum2(a int, b int) int {
	return a + b
}

// Anonymous Functions and Closures
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	// Multi return values
	quotient, remainder := divmod(7, 2)
	fmt.Println("quotient:", quotient, "remainder:", remainder)

	// Named return and Naked Return
	quotient, remainder = divmod2(7, 2)
	fmt.Println("quotient:", quotient, "remainder:", remainder)

	// Variadic Functions
	fmt.Println(sum(1, 2, 3, 4))

	// Functions as Values and Types
	fmt.Println(apply(sum2, 2, 3))

	// Anonymous Functions and Closures
	count := counter()
	fmt.Println(count())
	fmt.Println(count())
}
