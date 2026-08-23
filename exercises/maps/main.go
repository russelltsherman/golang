package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"blue":  "#0000FF",
		"green": "#00FF00",
		"white": "#FFFFFF",
	}

	// fmt.Println(colors)
	printMap((colors))
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)

	}
}
