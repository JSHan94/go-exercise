package main

import (
	"fmt"
)

func main() {
	// #1: declare the types
	type text struct {
		title string
		words int
	}

	type book struct {
		text
		isbn string

		title string // WARNING: conflict field
	}

	moby := book{
		text:  text{title: "moby dick", words: 206052},
		isbn:  "102030",
		title: "out",
	}

	moby.text.words = 1000
	moby.words++

	fmt.Printf("%s has %d words (isbn: %s)\n",
		moby.title, // equals to: moby.text.title
		moby.words, // equals to: moby.text.words
		moby.isbn)
	fmt.Printf(moby.title)
	fmt.Printf(moby.text.title)
}
