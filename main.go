package main

func main() {
	// Initialize a new deck type, which is
	// a slice of strings.
	cards := newDeck()

	// Specify two values that will receive the
	// two returned values from the function.
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
