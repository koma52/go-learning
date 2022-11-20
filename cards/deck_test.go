package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Two of Spades" {
		t.Errorf("Expected first card of Two of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Ace of Clubs" {
		t.Errorf("Expected last card of Ace of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktest")

	deck := newDeck()
	deck.saveToFile("_decktest")

	loadedDeck := newDeckFromFile("_decktest")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}

	os.Remove("_decktest")
}
