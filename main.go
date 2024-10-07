package main //executable type package
//package main requires function main, which will be called whenever we execute this file

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()
}
