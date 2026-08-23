// Here is where we explain the package.
//
// Some other stuff.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	values := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func newDeckFromFile(filename string) (deck, error) {
	bytes, error := os.ReadFile(filename)

	cards := deck(strings.Split(string(bytes), ","))

	return cards, error
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle() {
	// instanciate random number generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// iterate deck
	for i := range d {
		n := r.Intn(len(d) - 1)

		// swap values at indices i and n
		d[i], d[n] = d[n], d[i]
	}
}

func (d deck) toFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0660)
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
