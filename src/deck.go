package main

import "fmt"

// Create new type `deck` (slice of strings)
type deck []string

func makeDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardVals := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, val := range cardVals {
			cards = append(cards, val + " of " + suit)
		}
	}
	return cards
}

// Custom print method for `deck`
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deal 2x player hands from deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}