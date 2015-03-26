package uno

import (
	"cards/cards"
	"fmt"
)

type FuncCard struct {
	discardPile *cards.DiscardPile
	Name        string
	Color       Color
	Points      int
	isValidPlay func() bool
}

func NewFuncCard(color Color, name string, points int, discard *cards.DiscardPile) *FuncCard {
	c := &FuncCard{
		Name:        name,
		Color:       color,
		Points:      points,
		discardPile: discard,
	}
	return c
}

func (c FuncCard) String() string {
	if c.Color.String() == c.Name {
		return fmt.Sprintf("%s (%d pts)", c.Name, c.Points)
	}
	return fmt.Sprintf("%s %s (%d pts)", c.Color, c.Name, c.Points)
}

func (c FuncCard) Discard() {
	c.discardPile.Discard(c)
}

func NewWildCard(discard *cards.DiscardPile) cards.Card {
	return NewFuncCard(Wild, "Wild", 50, discard)
}

func NewDrawFourCard(discard *cards.DiscardPile) cards.Card {
	return NewFuncCard(Wild, "Draw Four", 50, discard)
}

func NewReverseCard(color Color, discard *cards.DiscardPile) cards.Card {
	return NewFuncCard(color, "Reverse", 20, discard)
}

func NewSkipCard(color Color, discard *cards.DiscardPile) cards.Card {
	return NewFuncCard(color, "Skip", 20, discard)
}

func NewDrawTwoCard(color Color, discard *cards.DiscardPile) cards.Card {
	return NewFuncCard(color, "Draw Two", 20, discard)
}
