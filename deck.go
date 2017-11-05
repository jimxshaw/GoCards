package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new deck type
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	// For unused variables like the i and j iterators in the
	// loops, replace them with the _ underscore to prevent
	// linting errors.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Go has support to return more than one value.
func deal(d deck, handSize int) (deck, deck) {
	// First slice contains values from the
	// start of the slice up to but not including
	// the handSize and second slice contains values
	// from the start of the handSize until the end.
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// Convert a deck in to a slice of strings.
	// Join the slice together separated by a comma.
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
