package main

import (
	"os"
	"testing"
)

const serializedDeck = "Ace of Clubs,Two of Clubs,Three of Clubs,Four of Clubs,Ace of Diamonds,Two of Diamonds,Three of Diamonds,Four of Diamonds,Ace of Hearts,Two of Hearts,Three of Hearts,Four of Hearts,Ace of Spades,Two of Spades,Three of Spades,Four of Spades"
const testFile = "_decktesting"

func TestNewDeckLength(t *testing.T) {
	d := newDeck()
	actual := len(d)
	expected := 16
	if expected != actual {
		t.Errorf("Expected deck length of %v, but got %v", expected, actual)
	}
}

func TestNewDeckFirstCard(t *testing.T) {
	d := newDeck()
	actual := d[0]
	expected := "Ace of Clubs"
	if expected != actual {
		t.Errorf("Expected first card to be %v, but got %v", expected, actual)
	}
}

func TestNewDeckLastCard(t *testing.T) {
	d := newDeck()
	actual := d[len(d)-1]
	expected := "Four of Spades"
	if actual != expected {
		t.Errorf("Expected last card to be %v, but got %v", expected, actual)
	}
}

func TestShuffleFirstCard(t *testing.T) {
	d := newDeck()
	d.shuffle()

	actual := d[0]
	expected := "Ace of Clubs"
	if expected == actual {
		t.Errorf("Expected first card not to be %v, but got %v", expected, actual)
	}
}

func TestShuffleLastCard(t *testing.T) {
	d := newDeck()
	d.shuffle()

	actual := d[len(d)-1]
	expected := "Four of Spades"
	if expected == actual {
		t.Errorf("Expected last card not to be %v, but got %v", expected, actual)
	}
}

func TestNewDeckToFile(t *testing.T) {
	os.Remove(testFile)

	d := newDeck()
	d.toFile(testFile)

	bytes, _ := os.ReadFile(testFile)

	actual := string(bytes)
	expected := serializedDeck
	if expected != actual {
		t.Errorf("Expected saved deck to be %v, but got %v", expected, actual)
	}
	os.Remove(testFile)
}

func TestNewDeckFromFile(t *testing.T) {
	os.Remove(testFile)
	os.WriteFile(testFile, []byte(serializedDeck), 0660)

	d, _ := newDeckFromFile(testFile)
	actual := len(d)
	expected := 16
	if expected != actual {
		t.Errorf("Expected deck length of %v, but got %v", expected, actual)
	}

	os.Remove(testFile)
}
