package main

import "fmt"

// Create a new type "Deck" that is a Slice of type String
type deck []string

// Print out the value of each card in a deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
