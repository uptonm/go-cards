package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck of length 52, but got %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected deck[i] to be 'Ace of Spades', but got '%v'", d[0])
	}

	if d[len(d)-1] != "King of Spades" {
		t.Errorf("Expected deck[i] to be 'King of Spades', but got '%v'", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck of length 52, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
