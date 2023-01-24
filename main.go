package main

func main() {
	card := newDeck()

	hand, remainingCards := deal(card, 5)

	hand.print()
	remainingCards.print()

}
