package main

import (
	"fmt"
)

// this is a comment
func init() {
	fmt.Println("Initializing application data")
}

func main() {
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
}
