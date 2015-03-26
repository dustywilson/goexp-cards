package uno

type Color int

const (
	Wild Color = iota
	Blue
	Green
	Red
	Yellow
)

var colors = map[Color]string{
	Wild:   "Wild",
	Blue:   "Blue",
	Green:  "Green",
	Red:    "Red",
	Yellow: "Yellow",
}

func (c Color) String() string {
	return colors[c]
}
