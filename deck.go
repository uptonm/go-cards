package main

import (
	"fmt"
	"strings"
)

// Create a new type "Deck" that is a Slice of type String
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// Index is replaced with underscore to tell Go that we don't want to use it
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Print out the value of each card in a deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Return a hand (deck of `handSize` cards) and the remaining deck (deck of d.size - `handSize` cards)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] // Return everything upto `handSize` and everything after it
}

// Convert deck to string joined by commas
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
