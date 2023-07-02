package main

import "fmt"

func process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover panic:", r)
		}
	}()

	fmt.Println("processing...")
	panic("panic occur!")
	fmt.Println("NOT CALLED!")
}

func main() {
	fmt.Println("start program")
	process()
	fmt.Println("end program")
}
