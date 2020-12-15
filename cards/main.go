package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" declaration
	// card := "Ace of Spades" declaration
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
