package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

func (eb englishBot) getGreeting() string {
	return "Hello there"
}

type spanishBot struct{}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
