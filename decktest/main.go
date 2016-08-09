// decktest is just a place for me to figure out how to draw a single card
// from a 1- or 9-stack deck for blackjack purposes (since, as of now,
// my blackjack program just draws from an "infinite" deck, with no control
// over how many of what suit/type of card gets drawn). This should make
// it closer to an actual, casino-style game of blackjack.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type shuffler interface {
	Len() int
	Swap(i, j int)
}

func shuffle(s shuffler) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < s.Len(); i++ {
		j := r.Intn(s.Len()-i)
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

func main() {
	deck := make([]string, 0)
	c := map[int]string{
		1:"Ace",
		2:"Two",
		3:"Three",
		4:"Four",
		5:"Five",
		6:"Six",
		7:"Seven",
		8:"Eight",
		9:"Nine",
		10:"Ten",
		11:"Jack",
		12:"Queen",
		13:"King",
	}
	suit := [4]string{"Clubs", "Diamonds", "Hearts", "Spades"}

	for j := 0; j < 4; j++ {
		for i := 1; i < 14; i++ {
			deck = append(deck, fmt.Sprint(c[i], " of ", suit[j]))
		}
	}


	fmt.Println("Unshuffled deck:")
	for x := range deck {
		fmt.Printf("%d: %s\n", x+1, deck[x])
	}

	shuffle(cardSlice(deck))
	shuffle(cardSlice(deck))
	shuffle(cardSlice(deck))

	fmt.Println("Shuffled deck:")
	for x := range deck {
		fmt.Printf("%d: %s\n", x+1, deck[x])
	}

	fmt.Println("Drawing from top of deck...")
	topcard := deck[:1]
	deck = deck[1:]
	fmt.Println("Your card...:", topcard[0])
	fmt.Println("There are now", len(deck), "cards in the deck.")
}
