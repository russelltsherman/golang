package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func echoworker(in, out chan int) {
	for {
		n := <-in
		time.Sleep(time.Duration(rand.Intn(3000) * int(time.Millisecond)))
		out <- n
	}
}

func produce(ch chan<- int) {
	i := 0
	for {
		fmt.Printf("-> send job: %d \n", i)
		ch <- i
		i++
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)
	for i := 0; i < 10; i++ {
		go echoworker(in, out)
	}
	go produce(in)
	for n := range out {
		fmt.Printf("<- Recv job: %d \n", n)
	}
}
