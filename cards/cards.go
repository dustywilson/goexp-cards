package cards

type Card interface {
	String() string
	Discard()
}
