package main //executable type package
//package main requires function main, which will be called whenever we execute this file
import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
}
