package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Testing length
	if len(d) != 52 {
		t.Errorf("Expected deck of length 52, but got %d", len(d))
	}

	// Testing first card
	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card as Ace of Heards, but got %s", d[0])
	}

	// Testing second card
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card as King of Clubs, but got %s", d[0])
	}

	// Testing for duplicates
	for i, iVal := range d {
		for j, jVal := range d[i+1:] {
			if iVal == jVal {
				t.Errorf("Found duplicate value %s at %d and %d", iVal, i, j)
			}
		}
	}

}

func TestSaveToFileAndReadDeckFromFile(t *testing.T) {
	testFile := "_testfile"
	d := newDeck()
	d.saveToFile(testFile)
	deckFromFile := newDeckFromFile(testFile)
	os.Remove(testFile)
	if len(deckFromFile) != len(d) {
		t.Errorf("Loaded deck and source deck length does not match")
	}

}
