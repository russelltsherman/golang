package main

import (
	"context"
	"fmt"
	"time"
)

func init() {
	fmt.Println("Initializing application data")
}

func main() {
	ctx := context.Background()
	done := ctx.Done()

	for i := 0; ; i++ {
		select {
		case <-done:
			return
		case <-time.After(time.Second):
			fmt.Println("tick: ", i)
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
}
