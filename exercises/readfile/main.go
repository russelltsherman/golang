package main

import (
	"fmt"
	"io"
	"os"
)

func handle_err(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	filename := os.Args[1]
	fmt.Println(filename)

	file, err := os.Open(filename)
	handle_err(err)
	io.Copy(os.Stdout, file)

	// bs, err := os.ReadFile(filename)
	// handle_err(err)
	// fmt.Println(string(bs))
}
