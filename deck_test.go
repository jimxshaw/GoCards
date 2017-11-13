package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	totalCards := 52

	// 4 card suits multiply by 13 cards
	// equals total cards.
	if len(d) != totalCards {
		t.Errorf("Expected deck length of %v but got %v", totalCards, len(d))
	}

	// Make sure the first card in the deck is
	// what we want it to be.
	if d[0] != "A of Spades" {
		t.Errorf("Expected first card to be A of Spades but instead got %v", d[0])
	}

	if d[len(d)-1] != "K of Clubs" {
		t.Errorf("Expected last card to be K of Clubs but instead got %v", d[len(d)-1])
	}
}
