package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	count := 52
	if len(d) != count {
		t.Errorf("Expected deck length of %v, but got %v", count, len(d))
	}

	first := "Ace of Clubs"
	if d[0] != first {
		t.Errorf("Expected first card to be %v, but got %v", first, d[0])
	}

	last := "King of Spades"
	if d[51] != last {
		t.Errorf("Expected last card to be %v, but got %v", last, d[0])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	testfile := "_decktest"
	os.Remove(testfile)

	d := newDeck()
	d.saveToFile(testfile)

	loaded := newDeckFromFile(testfile)

	count := 52
	if len(loaded) != count {
		t.Errorf("Expected deck length of %v, but got %v", count, len(loaded))
	}

	os.Remove(testfile)
}
