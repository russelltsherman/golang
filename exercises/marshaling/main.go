package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name     string `json:"username"`
	password string
}

func main() {
	fmt.Println("Hello")

	users := []user{
		{"inanc", "1234"},
		{"god", "42"},
		{"devil", "66"},
	}
	fmt.Println(users)

	out, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}
