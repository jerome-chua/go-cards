package main

func main() {	
	// Slice of cards
	cards := makeDeck() 
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
}

