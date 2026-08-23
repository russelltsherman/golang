package main

import (
	"fmt"
	"os"
)

func handle_err(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	cards := newDeck()
	cards.shuffle(3)
	cards.print()

	// hand, remainingCards := deal(cards, 5)
	// cards = remainingCards
	// hand.print()
}
