package main

import "fmt"

func main() {
	cards := []string{"Two of Hearts", newCard()}
	cards = append(cards, "Jack of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Seven of Clubs"
}
