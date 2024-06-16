package main

func main() {

	// fileName := "my_cards"

	cards := newDeck()

	cards.shuffleDeck()
	cards.printDeck()

}
