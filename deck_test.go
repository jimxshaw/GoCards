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

}
