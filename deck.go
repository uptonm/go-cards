package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type card struct {
	suit  string
	value string
}

// Create a new type "Deck" that is a Slice of type String
type deck []card

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// Index is replaced with underscore to tell Go that we don't want to use it
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, card{suit, value})
		}
	}

	return cards
}

// Print out the value of each card in a deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card.value+" of "+card.suit)
	}
}

// Return a hand (deck of `handSize` cards) and the remaining deck (deck of len(d) - `handSize` cards)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] // Return everything upto `handSize` and everything after it
}

// Convert deck to string joined by commas
func (d deck) toString() string {
	arr := []string{}
	for _, card := range d {
		arr = append(arr, card.value+" of "+card.suit)
	}
	return strings.Join(arr, ",")
}

// Save deck to file for repeated usage, returns error if thrown
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Reads byte string from file, splits at "," to create []string, then converts to deck type
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Err:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	d := deck{}
	for _, str := range s {
		d = append(d, cardFromString(str))
	}
	return d
}

func cardFromString(s string) card {
	values := strings.Split(s, " of ")
	return card{values[0], values[1]}
}

// Loops through deck, moves deck[i] to random position in deck
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
