package main

import "fmt"

func main() {
	studentGrades := make(map[string]map[string]int)

	studentGrades["Alice"] = map[string]int{"Math": 90, "Science": 95}
	studentGrades["Bob"] = map[string]int{"Math": 85, "Science": 92}

	fmt.Println(studentGrades["Alice"]["Math"])
	fmt.Println(studentGrades["Bob"]["Science"])
}
