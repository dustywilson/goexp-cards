package cards

import "sync"

type DiscardPile struct {
	cards []Card
	sync.RWMutex
}

func NewDiscardPile() *DiscardPile {
	return &DiscardPile{
		cards: make([]Card, 0),
	}
}

func (dp *DiscardPile) Discard(c Card) {
	dp.Lock()
	defer dp.Unlock()
	dp.cards = append(dp.cards, c)
}

func (dp *DiscardPile) ShuffleInto(d Deck) {
	dp.Lock()
	defer dp.Unlock()

	var lastCard Card
	lastCard, dp.cards = dp.cards[len(dp.cards)-1], dp.cards[:len(dp.cards)-1]
	d.ShuffleCards(dp.cards)
	dp.cards = []Card{lastCard}

	// UNO rules state that the very last discarded card starts the new discard pile
	// which is why we keep "lastCard" in the set when wiping out the array above.
	// This is so the game can continue with the same number/color/action that it
	// was at before shuffling the discard pile.
	// FIXME: the above is only true for times that the deck runs out completely... do we care?
	// FIXME: UNO-specific details do NOT belong in the "cards" package!
}

func (dp *DiscardPile) TopCard() Card {
	return dp.cards[len(dp.cards)-1]
}
