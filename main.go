package main

func main() {
	cards := deck{newCard(), newCard()}
	cards = append(cards, newCard())

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
