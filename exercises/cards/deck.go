package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new data type of 'deck'
// a deck exists as a slice of strings
type deck []string

func newDeck() deck {
	d := deck{}
	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cards := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, s := range suits {
		for _, c := range cards {
			d = append(d, c+" of "+s)
		}
	}
	return d
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	handle_err(err)
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle(iterations int) {
	for i := 0; i < iterations; i++ {
		src := rand.NewSource(time.Now().UnixNano())
		r := rand.New(src)

		for idx := range d {
			newPosition := r.Intn(len(d) - 1)
			d[idx], d[newPosition] = d[newPosition], d[idx]
		}
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}
