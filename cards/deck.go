package main

import (
	"fmt"
	"os"
	"strings"
)

// Create a new type of 'deck'
// This is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			newCard := value + " of " + suit
			cards = append(cards, newCard)
		}
	}

	return cards
}

func (d deck) printDeck() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func dealCards(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
