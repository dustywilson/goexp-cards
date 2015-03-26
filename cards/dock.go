package cards

type Deck interface {
	Draw() Card
	ShuffleCards([]Card)
}
