package main

import (
	"os"
	"testing"
)

func TestCreateCards(t *testing.T) {
	d := createCards()

	if len(d) != 12 {
		t.Errorf("Expected deck length is 16, but got %v", len(d))
	}

	if d[0] != "Ace for Spades" {
		t.Errorf("Expected first deck Ace for Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four for Diamonds" {
		t.Errorf("Expected last deck Four for Diamonds, but got %v", d[len(d)-1])
	}
}

func TestCreateNewDeckAndDeleteDeck(t *testing.T) {
	os.Remove("__testingDeck")

	deck := createCards()
	deck.saveToFile("__testingDeck")

	loadedDeck := cardsFromFile("__testingDeck")

	if len(loadedDeck) != 12 {
		t.Errorf("Expected deck length is 12, but got %v", len(loadedDeck))
	}

	os.Remove("__testingDeck")
}
