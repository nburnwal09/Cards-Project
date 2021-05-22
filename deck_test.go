package main

import (
	"os"
	"testing"
)

//below function will be automatically called by go test runner with argument t which is of type *testing.T
func TestNewDeck(t *testing.T) { //T-->caps
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 20 but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("expected first card to be ace of cards but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("expected last card to be four of clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
