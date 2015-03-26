package uno

import (
	"cards/cards"
	"sync"
)

type Deck struct {
	cards []cards.Card
	sync.RWMutex
}

func NewDeck(discard *cards.DiscardPile) *Deck {
	deck := Deck{
		cards: []cards.Card{
			NewDrawFourCard(discard),
			NewDrawFourCard(discard),
			NewDrawFourCard(discard),
			NewDrawFourCard(discard),
			NewWildCard(discard),
			NewWildCard(discard),
			NewWildCard(discard),
			NewWildCard(discard),
			NewReverseCard(Blue, discard),
			NewReverseCard(Blue, discard),
			NewReverseCard(Green, discard),
			NewReverseCard(Green, discard),
			NewReverseCard(Red, discard),
			NewReverseCard(Red, discard),
			NewReverseCard(Yellow, discard),
			NewReverseCard(Yellow, discard),
			NewSkipCard(Blue, discard),
			NewSkipCard(Blue, discard),
			NewSkipCard(Green, discard),
			NewSkipCard(Green, discard),
			NewSkipCard(Red, discard),
			NewSkipCard(Red, discard),
			NewSkipCard(Yellow, discard),
			NewSkipCard(Yellow, discard),
			NewDrawTwoCard(Blue, discard),
			NewDrawTwoCard(Blue, discard),
			NewDrawTwoCard(Green, discard),
			NewDrawTwoCard(Green, discard),
			NewDrawTwoCard(Red, discard),
			NewDrawTwoCard(Red, discard),
			NewDrawTwoCard(Yellow, discard),
			NewDrawTwoCard(Yellow, discard),
			NewNumberCard(Blue, 0, discard),
			NewNumberCard(Blue, 0, discard),
			NewNumberCard(Blue, 1, discard),
			NewNumberCard(Blue, 1, discard),
			NewNumberCard(Blue, 2, discard),
			NewNumberCard(Blue, 2, discard),
			NewNumberCard(Blue, 3, discard),
			NewNumberCard(Blue, 3, discard),
			NewNumberCard(Blue, 4, discard),
			NewNumberCard(Blue, 4, discard),
			NewNumberCard(Blue, 5, discard),
			NewNumberCard(Blue, 5, discard),
			NewNumberCard(Blue, 6, discard),
			NewNumberCard(Blue, 6, discard),
			NewNumberCard(Blue, 7, discard),
			NewNumberCard(Blue, 7, discard),
			NewNumberCard(Blue, 8, discard),
			NewNumberCard(Blue, 8, discard),
			NewNumberCard(Blue, 9, discard),
			NewNumberCard(Blue, 9, discard),
			NewNumberCard(Green, 0, discard),
			NewNumberCard(Green, 0, discard),
			NewNumberCard(Green, 1, discard),
			NewNumberCard(Green, 1, discard),
			NewNumberCard(Green, 2, discard),
			NewNumberCard(Green, 2, discard),
			NewNumberCard(Green, 3, discard),
			NewNumberCard(Green, 3, discard),
			NewNumberCard(Green, 4, discard),
			NewNumberCard(Green, 4, discard),
			NewNumberCard(Green, 5, discard),
			NewNumberCard(Green, 5, discard),
			NewNumberCard(Green, 6, discard),
			NewNumberCard(Green, 6, discard),
			NewNumberCard(Green, 7, discard),
			NewNumberCard(Green, 7, discard),
			NewNumberCard(Green, 8, discard),
			NewNumberCard(Green, 8, discard),
			NewNumberCard(Green, 9, discard),
			NewNumberCard(Green, 9, discard),
			NewNumberCard(Red, 0, discard),
			NewNumberCard(Red, 0, discard),
			NewNumberCard(Red, 1, discard),
			NewNumberCard(Red, 1, discard),
			NewNumberCard(Red, 2, discard),
			NewNumberCard(Red, 2, discard),
			NewNumberCard(Red, 3, discard),
			NewNumberCard(Red, 3, discard),
			NewNumberCard(Red, 4, discard),
			NewNumberCard(Red, 4, discard),
			NewNumberCard(Red, 5, discard),
			NewNumberCard(Red, 5, discard),
			NewNumberCard(Red, 6, discard),
			NewNumberCard(Red, 6, discard),
			NewNumberCard(Red, 7, discard),
			NewNumberCard(Red, 7, discard),
			NewNumberCard(Red, 8, discard),
			NewNumberCard(Red, 8, discard),
			NewNumberCard(Red, 9, discard),
			NewNumberCard(Red, 9, discard),
			NewNumberCard(Yellow, 0, discard),
			NewNumberCard(Yellow, 0, discard),
			NewNumberCard(Yellow, 1, discard),
			NewNumberCard(Yellow, 1, discard),
			NewNumberCard(Yellow, 2, discard),
			NewNumberCard(Yellow, 2, discard),
			NewNumberCard(Yellow, 3, discard),
			NewNumberCard(Yellow, 3, discard),
			NewNumberCard(Yellow, 4, discard),
			NewNumberCard(Yellow, 4, discard),
			NewNumberCard(Yellow, 5, discard),
			NewNumberCard(Yellow, 5, discard),
			NewNumberCard(Yellow, 6, discard),
			NewNumberCard(Yellow, 6, discard),
			NewNumberCard(Yellow, 7, discard),
			NewNumberCard(Yellow, 7, discard),
			NewNumberCard(Yellow, 8, discard),
			NewNumberCard(Yellow, 8, discard),
			NewNumberCard(Yellow, 9, discard),
			NewNumberCard(Yellow, 9, discard),
		},
	}
	return &deck
}

func (d *Deck) Draw() (c cards.Card) {
	d.Lock()
	defer d.Unlock()
	if len(d.cards) == 0 {
		return nil
	}
	c, d.cards = d.cards[0], d.cards[1:]
	return c
}

// ShuffleCards is used to shuffle the deck along with some cards that are
// passed into this method.  It's really intended to be used automatically from
// the cards.DiscardPile.ShuffleInto(cards.Deck) method.
func (d *Deck) ShuffleCards(cs []cards.Card) {
	d.Lock()
	defer d.Unlock()
	d.cards = append(d.cards, cs...)
	cards.Shuffle(d.cards)
}

// Shuffle only shuffles what's currently in the deck and doesn't include any
// other cards, not even from the discard deck.
func (d *Deck) Shuffle() {
	d.ShuffleCards(nil)
}
