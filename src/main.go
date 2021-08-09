package main

func main() {	
	// Slice of cards
	cards := deck{newCard(), newCard()} 
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Ace of Hearts"
}