package main

func main() {
	// Initialize a new deck type, which is
	// a slice of strings.
	cards := deck{"Two of Hearts", newCard()}

	// Append a card to our cards variable and then
	// re-assign it back to our cards variable.
	// The append function does not mutate the
	// origin slice but will return a new slice.
	cards = append(cards, "Jack of Spades")

	cards.print()
}

func newCard() string {
	return "Seven of Clubs"
}
