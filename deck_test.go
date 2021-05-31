package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %d", len(d))
	}
	if d[0] != "Ace of Diamonds" {
		t.Errorf("Expected first card of Ace of Diamonds, but got %v", d[0])
	}
	if d[len(d)-1] != "four of Club" {
		t.Errorf("Expected first card of four of Club, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedFile := newDeckFromFile("_decktesting")
	if len(loadedFile) != 16 {
		t.Errorf("Expected 16 cards from the deck, but got %d", len(loadedFile))
	}
	os.Remove("_decktesting")
}
