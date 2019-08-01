package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())

	// hand, remainingCards := deal(cards, 3)

	// hand.print()
	// remainingCards.print()
}
