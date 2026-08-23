package main

import (
	"fmt"
)

func main() {

	numbs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range numbs {
		if num%2 == 0 {
			fmt.Println(num, "even")
		} else {
			fmt.Println(num, "odd")
		}
	}
}
