package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		// Option #1 - Log the error and return a call to newDeck().

		// Option #2 - Log the error and entirely quit the program.
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Convert byte slice into a string, split that string by comma
	// into a string slice and the convert it into a deck.
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	// A new rand object needs a source to generate a seed value.
	// We pass the current time in int64 format to generate that
	// new source to be used to generate a new rand object.
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// Generate a random int from the beginning of the deck
		// slice to however long the deck slice is.
		newPosition := r.Intn(len(d) - 1)

		// Within our deck slice, take the value at newPosition
		// and swap it with the value at i. Then take the value at
		// i and swap it with the value at newPosition.
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
