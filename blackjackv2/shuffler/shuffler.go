// Deck returns a []string populated with a shuffled 52-card deck.
// Can be used for any number of card games.
// Author: Todd S.
package shuffler

import (
	"math/rand"
	"time"
)

type shuffler interface {
	Len() int
	Swap(i, j int)
}

func shuffle(s shuffler) {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < s.Len(); i++ {
		j := r.Intn(s.Len() - i)
		s.Swap(i, j)
	}
}

type cardSlice []string

func (s cardSlice) Len() int {
	return len(s)
}

func (s cardSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func ShuffleDeck(deck []string) []string {
	for i := 0; i < 4; i++ {
		shuffle(cardSlice(deck))
	}

	return deck
}
