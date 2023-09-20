package main

import (
	"os"
	"testing"
)

//t.Errorf("Expected deck length of 16, but got %v", len(d))

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSavetoFileAndNewDeckFromFile(t *testing.T) {
	const FILENAME = "_decktesting"
	os.Remove(FILENAME)

	deck := newDeck()
	deck.saveToFile(FILENAME)
	loadedDeck := newDeckFromFile(FILENAME)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove(FILENAME)
}
