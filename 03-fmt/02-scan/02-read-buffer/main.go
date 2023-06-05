package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a int

	// fmt.Scan
	fmt.Print("Enter a number using fmt.Scan: ")
	fmt.Scan(&a)
	fmt.Println("You entered:", a)

	// stdin.ReadString
	fmt.Print("Enter a string using stdin.ReadString: ")
	reader := bufio.NewReader(os.Stdin)
	b, _ := reader.ReadString('\n')
	fmt.Println("You entered:", b)
}
