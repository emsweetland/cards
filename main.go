package main //executable type package
//package main requires function main, which will be called whenever we execute this file

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

}