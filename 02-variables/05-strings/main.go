package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, World!"

	fmt.Println(len(s))                                // "13"
	fmt.Println(strings.Contains("test", "es"))        // "true"
	fmt.Println(strings.Count("test", "t"))            // "2"
	fmt.Println(strings.HasPrefix("test", "te"))       // "true"
	fmt.Println(strings.HasSuffix("test", "st"))       // "true"
	fmt.Println(strings.Index("test", "e"))            // "1"
	fmt.Println(strings.Join([]string{"a", "b"}, "-")) // "a-b"
	fmt.Println(strings.Repeat("a", 5))                // "aaaaa"
	fmt.Println(strings.Split("a-b", "-"))             // ["a", "b"]
}
