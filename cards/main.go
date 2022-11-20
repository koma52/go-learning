package main

import "fmt"

func main() {
	// var card string = "Ace of spades"
	// card := newCard()
	cards := []string{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	fmt.Println(cards)
	// fmt.Println(card)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
