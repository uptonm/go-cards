package main

func main() {
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards.txt")
	// hand, remainingCards := deal(cards, 3)

	// hand.print()
	// remainingCards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
