package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
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
		fmt.Println((i + 1), card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) (d deck) {
	bs, err := os.ReadFile(fileName)
	if err != nil {
		log.Println("ERROR: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func dealCards(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func generateRandSource() int64 {
	t := time.Now()
	return t.UnixNano()
}

func (d deck) shuffleDeck() {

	// Using the built in function
	// rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })

	// Custom logic
	r := rand.New(rand.NewSource(generateRandSource()))

	for i := range d {

		r := r.Intn(len(d))
		d[i], d[r] = d[r], d[i]
	}

}
