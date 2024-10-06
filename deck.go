package main

import (
	"fmt"
	"strings"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// newDeck will return value of type deck
// does not need a reciever because the purpose is to create a new deck, we do not have one yet
func newDeck() deck {
	// create an empty deck
	cards := deck{}
	// create two slices for card suits and card values
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// create each card
	// _ used instead of i because we are not going to use the index this time
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// add each card to empty deck (cards)
			cards = append(cards, value+" of "+suit)
		}
	}
	// return slice of cards we've created
	return cards
}

// reciever of deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// using multiple arguments inside of function call, return two values of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// helper function to recieve a deck and return string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}