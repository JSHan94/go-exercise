package main

import "fmt"

// Basic function
func sum(a int, b int) int {
	return a + b
}

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
func sum2(nums ...int) int {
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

// Anonymous Functions and Closures
func counter() func() int {
	count := 0
	return func() int {
		// count 항상 기억되고 capture되어 있다
		count++
		return count
	}
}

func main() {
	// Basic Function
	// fmt.Println(apply(sum, 2, 3))

	// Multi return values
	// quotient, remainder := divmod(7, 2)
	// fmt.Println("quotient:", quotient, "remainder:", remainder)

	// Named return and Naked Return
	// quotient, remainder = divmod2(7, 2)
	// fmt.Println("quotient:", quotient, "remainder:", remainder)

	// Variadic Functions
	// fmt.Println(sum2(1, 2, 3, 4, 5, 6, 7))

	// Anonymous Functions and Closures
	countee := counter() // count
	fmt.Println(countee())
	fmt.Println(countee())
	fmt.Println(countee())
	fmt.Println(countee())
	fmt.Println(countee())
	fmt.Println(countee())

	// 데이터 은닉 (캡슐화)
	// 런타임 함수를 변수로
	// 값을 계속 기억 -> 로컬 변수의 상태를 유지하는데 유용하다
}
