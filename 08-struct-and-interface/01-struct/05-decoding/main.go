package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type user struct {
	Name        string          `json:"username"`
	Permissions map[string]bool `json:"perms"`
}

func main() {
	data, err := ioutil.ReadFile("./users.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var users []user
	if err := json.Unmarshal(data, &users); err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Print("+ ", user.Name)

		switch p := user.Permissions; {
		case p == nil:
			fmt.Print(" has no power.")
		case p["admin"]:
			fmt.Print(" is an admin.")
		case p["write"]:
			fmt.Print(" can write.")
		}
		fmt.Println()
	}
}
