package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// no reciever, getting access to a deck we don't have yet
func newDeckFromFile(filename string) deck {
	// ReadFile will return a byte slice, and an error
	bs, err := os.ReadFile(filename)
	//check if error was returned
	if err != nil {
		// option #1 - log the error, return call to newDeck()
		// option #1 - log the error, quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// s for slice of strings, slice, etc
	s := strings.Split(string(bs), ",")
	// type conversion from slice of string to deck
	return deck(s)
}

// recieves a deck to shuffle
func (d deck) shuffle() {
	// using time to seed our random generator
	// this is no longer necessary, as of Go 1.20
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// index in the range of deck
	for i := range d {
		// random integer within the length of deck -1
		newPosition := r.Intn(len(d) - 1)
		// swap elements at their positions
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
