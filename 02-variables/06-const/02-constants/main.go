package main

import "fmt"

func main() {
	const (
		a = iota // a == 0
		b        // b == 1
		c        // c == 2
	)

	fmt.Println(a, b, c)
}
