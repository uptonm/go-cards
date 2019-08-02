package main

func main() {
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	// hand, remainingCards := deal(cards, 3)

	// hand.print()
	// remainingCards.print()
	cards := newDeckFromFile("my_cards")
	cards.print()
}
