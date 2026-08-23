package main

import "fmt"

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	t := triangle{base: 3, height: 3}
	s := square{sideLength: 3}

	printArea(t)
	printArea(s)
}
