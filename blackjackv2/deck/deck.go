// Deck returns a []string populated with a shuffled 52-card deck.
// Can be used for any number of card games.
// Author: Todd S.
package deck

import (
	"fmt"
)

func Deck() []string {
	var deck []string
	c := map[int]string{
		1:  "Ace",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
		11: "Jack",
		12: "Queen",
		13: "King",
	}
	suit := [4]string{"Clubs", "Diamonds", "Hearts", "Spades"}

	for j := 0; j < 4; j++ {
		for i := 1; i < 14; i++ {
			deck = append(deck, fmt.Sprint(c[i], " of ", suit[j]))
		}
	}
	return deck
}
