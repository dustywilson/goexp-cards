package cards

import "math/rand"

func Shuffle(cs []Card) {
	l := len(cs)
	for i := 0; i < l; i++ {
		j := rand.Intn(l-i) + i
		cs[j], cs[i] = cs[i], cs[j]
	}
}
