package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// create a deck and save it
	cards := newDeck()

	if len(cards) != 16 {
		t.Errorf("expected length 16 YOU DUMB FUCK but got %v", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("expected Ace of Spades YOU DUMB FUCK but got %v", cards[0])
	}

	if cards[len(cards)-1] != "Four of Clubs" {
		t.Errorf("expected Four of Clubs YOU DUMB FUCK but got %v", cards[4])
	}
}

func TestSaveToDeckAndNewDeckFromFileFunction(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("expected length 16 YOU DUMB FUCK but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
