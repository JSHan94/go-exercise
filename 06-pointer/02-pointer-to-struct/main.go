package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// pointer receiver
func (p *Person) SetName(name string) {
	p.Name = name
}

func main() {
	// Create a Person struct.
	p := &Person{Name: "John Doe", Age: 30}

	fmt.Println(p.Name) // John Doe
	fmt.Println(p.Age)  // 30

	// Call the SetName() function, passing in a pointer to the struct.
	p.SetName("Jane Doe")

	// Print the name of the struct.
	fmt.Println(p.Name) // Jane Doe
}
