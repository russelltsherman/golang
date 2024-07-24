package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func producer(ch chan<- int, name string) {
	for {
		// sleep for a random time
		time.Sleep(time.Duration(rand.Intn(3000)))

		// generate a random number
		n := rand.Intn(100)

		// send a message
		fmt.Printf("Channel %s -> %d \n", name, n)
		ch <- n
	}
}

func consumer(ch <-chan int) {
	for n := range ch {
		fmt.Printf("<- %d \n", n)
	}
}

func fanIn(chA, chB <-chan int, chC chan<- int) {
	var n int
	for {
		select {
		case n = <-chA:
			chC <- n
		case n = <-chB:
			chC <- n
		}
	}
}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go producer(chA, "A")
	go producer(chB, "B")
	go consumer(chC)

	fanIn(chA, chB, chC)
}
