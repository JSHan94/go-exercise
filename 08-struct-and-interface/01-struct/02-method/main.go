package main

import (
	"fmt"
)

type Rectangle struct {
	width  int
	height int
}

func (r Rectangle) calculateArea() int {
	return r.width * r.height
}

func main() {
	rect := Rectangle{
		width:  10,
		height: 5,
	}

	area := rect.calculateArea()

	fmt.Println("Area of the rectangle:", area)
}
