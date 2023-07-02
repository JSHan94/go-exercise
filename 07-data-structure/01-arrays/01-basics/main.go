package main

import "fmt"

func main() {
	// var (
	// 	myAge  byte
	// 	herAge byte

	// 	// its length      : 2
	// 	// its element type: byte
	// 	ages [2]byte

	// 	// its length      : 5
	// 	// its element type: string
	// 	tags [5]string

	// 	// don't occupy any memory space
	// 	// its length      : 0
	// 	// its element type: byte
	// 	zero [0]byte

	// 	// its length      : 3
	// 	// its element type: byte
	// 	agesExp [2 + 1]byte
	// )

	// fmt.Printf("myAge             : %d\n", myAge)
	// fmt.Printf("herAge            : %d\n", herAge)

	// // #v verb prints the array's length, element type and its elements
	// fmt.Printf("ages              : %#v\n", ages)
	// fmt.Printf("tags              : %#v\n", tags)
	// fmt.Printf("zero              : %#v\n", zero)
	// fmt.Printf("agesExp           : %#v\n", agesExp)

	{
		var ages [2]int

		fmt.Println()
		fmt.Printf("ages              : %#v\n", ages)
		fmt.Printf("ages's type       : %T\n", ages)

		fmt.Println("len(ages)         :", len(ages))
		fmt.Println("ages[0]           :", ages[0])
		fmt.Println("ages[1]           :", ages[1])
		fmt.Println("ages[len(ages)-1] :", ages[len(ages)-1])

		// WRONG:
		// fmt.Println("ages[-1]          :", ages[-1])
		// fmt.Println("ages[2]           :", ages[2])
		// fmt.Println("ages[len(ages)]   :", ages[len(ages)])

		ages[0] = 6
		ages[1] -= 3

		// WRONG:
		// ages[0] = "Can I?"

		fmt.Println("ages[0]           :", ages[0])
		fmt.Println("ages[1]           :", ages[1])

		ages[0] *= ages[1]
		fmt.Println("ages[0]           :", ages[0])
		fmt.Println("ages[1]           :", ages[1])
	}
}
