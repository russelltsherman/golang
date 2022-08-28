package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name string `json:"username"`
}

func main() {

	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	var users []user
	if err := json.Unmarshal(input, &users); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(users)

}
