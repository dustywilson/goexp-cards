package main

import (
	"cards/cards"
	"cards/uno"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	discard := cards.NewDiscardPile()
	deck := uno.NewDeck(discard)

	deck.Shuffle()
	deck.Draw().Discard() // UNO rules say we must discard the top card from the deck on a new game

	fmt.Printf("[Top of DiscardPile] %s\n", discard.TopCard())

	fmt.Println("Deck One:")
	for {
		c := deck.Draw()
		if c == nil {
			break
		}
		fmt.Printf("\tDraw: %s\n", c)
		c.Discard()
	}

	discard.ShuffleInto(deck)
	fmt.Printf("[Top of DiscardPile] %s\n", discard.TopCard())

	fmt.Println("Deck Two:")
	for {
		c := deck.Draw()
		if c == nil {
			break
		}
		fmt.Printf("\tDraw: %s\n", c)
		c.Discard()
	}
}
