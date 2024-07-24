package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func process2(ch chan int) {
	n := rand.Intn(3000)
	time.Sleep(time.Duration(n) * time.Millisecond)
	ch <- n
}

func main() {
	// create channel
	ch := make(chan int)

	// spawn process
	go process2(ch)

	fmt.Println("Waiting for process")
	n := <-ch
	fmt.Printf("Process took %dms\n", n)
}
