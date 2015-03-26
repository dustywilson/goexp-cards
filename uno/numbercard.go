package uno

import (
	"cards/cards"
	"fmt"
)

type NumberCard struct {
	discardPile *cards.DiscardPile
	Color       Color
	Number      Number
	Points      int
}

func NewNumberCard(color Color, number int, discard *cards.DiscardPile) *NumberCard {
	c := &NumberCard{
		Color:       color,
		Number:      Number(number),
		Points:      number,
		discardPile: discard,
	}
	return c
}

func (c NumberCard) String() string {
	return fmt.Sprintf("%s %d (%d pts)", c.Color, c.Number, c.Points)
}

func (c NumberCard) Discard() {
	c.discardPile.Discard(c)
}
