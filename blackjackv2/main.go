// blackjack draws from a single-stack deck.// TODO: Add option to pull from single OR 9-stack deck.
// Deals initial 2 cards, then asks user if they'd like to hit (default is yes) until
// user wins (score == 21; blackjack), busts (score > 21; bust), gets 5 cards, or chooses
// to stay.
// Author: Todd S.
package main

import (
	"fmt"
	"strings"

	"./deck"
	"./shuffler"
)

var shuffleMe = deck.Deck()
var d = shuffle(shuffleMe) // d=deck after the shuffle.

type Card struct {
	card  string
	value int
}

// shuffle takes a slice of strings (52 cards) and calls the imported
// shuffler.ShuffleDeck() with them (shuffleMe, renamed as d
// just for this function).
func shuffle(d []string) []string {
	fmt.Println("Shuffling deck...")
	shuffler.ShuffleDeck(d)
	return d
}

func askYn(q string) bool {
	var a string
	fmt.Printf(q)
	_, err := fmt.Scanln(&a)
	strings.ToLower(a)
	switch {
	case strings.HasPrefix("yes", a):
		return true
	case strings.HasPrefix("no", a):
		return false
	case fmt.Sprint(err) == "unexpected newline":
		return true
	default:
		fmt.Println("Just a y or n will do.")
		return askYn(q)
	}
}

// cardEval (card evaluator) applies blackjack scoring rules
// to each card sent to it, returning a corresponding score.
func cardEval(c string) int {
	switch {
	case strings.HasPrefix(c, "Ace"):
		return 11
	case strings.HasPrefix(c, "Two"):
		return 2
	case strings.HasPrefix(c, "Three"):
		return 3
	case strings.HasPrefix(c, "Four"):
		return 4
	case strings.HasPrefix(c, "Five"):
		return 5
	case strings.HasPrefix(c, "Six"):
		return 6
	case strings.HasPrefix(c, "Seven"):
		return 7
	case strings.HasPrefix(c, "Eight"):
		return 8
	case strings.HasPrefix(c, "Nine"):
		return 9
	default:
		return 10
	}
}

// The aptly named 'draw' does just that, drawing/returning the 'top'
// card from the deck, after which deck equals deck, from the
// second card 'til the end (d = d[1:]).
func draw() (string, int) {
	topcard := d[0]
	d = d[1:]
	value := cardEval(topcard)
	return topcard, value
}

// handInit (hand-initializer) is called to get the first 2 cards
// of the hand for both dealer and player.
func handInit() ([]string, int, bool) {
	var highace bool
	hand := make([]string, 2)

	x, y := draw()
	c1 := Card{x, y}
	x, y = draw()
	c2 := Card{x, y}

	if strings.Contains(c1.card, "Ace") {
		highace = true
	}
	if strings.Contains(c2.card, "Ace") && highace == true {
		c2.value = 1
	}

	hand[0] = c1.card
	hand[1] = c2.card

	t := c1.value + c2.value

	return hand, t, highace
}

// play is where the player makes decisions about whether to stand or hit.
// function returns the total score (t), full hand (hand), and whether
// or not we're still holding a high ace (highace).
func play(t int, hand []string, highace bool) (int, []string, bool) {
	for t < 21 {
		if askYn("Would you like to hit? [Y/n]: ") {
			nc, nv := draw() // nc=new card; nv=new value.
			switch {
			case strings.Contains(nc, "Ace") && t+nv <= 21:
				highace = true
				t += nv
			case strings.Contains(nc, "Ace") && t+nv > 21:
				t++
			default:
				t += nv
			}
			hand = append(hand, nc)
			fmt.Println("*** Your next card:", nc)
		} else {
			break
		}
		if highace == true && t > 21 {
			fmt.Println("High Ace ('soft' hand) becomes Low Ace ('hard' hand)...")
			highace = false
			t -= 10
		}
		if t < 21 && len(hand) == 5 {
			fmt.Println("5-card hand... YOU WIN!! :D")
			break
		}
		switch {
		case t == 21:
			fmt.Println("BLACKJACK!! :D")
		case t > 21:
			fmt.Println("B U S T E D! :(")
		default:
			fmt.Println("Current score:", t)
		}
	}
	return t, hand, highace
}

func showHand(hand []string) {
	for i := range hand {
		fmt.Printf("%d. %s\n", i+1, hand[i])
	}
	fmt.Println()
}


func main() {
	dd := make([]string, 0) // Discarded Deck
	var handcount int
	for handcount < 5 {
		hand, t, highace := handInit()

		fmt.Println("*** Your Hand: ***")
		showHand(hand)

		fmt.Println("*** Dealer's Hand: ***")
		dhand, dt, _ := handInit()
		fmt.Printf("1. %s\n2. (Face Down)\n\n", dhand[0])

		fmt.Println("Score:", t)
		if t == 21 {
			fmt.Println("BLACKJACK!! :D")
		} else {
			t, hand, highace = play(t, hand, highace)
		}

		fmt.Println(" - Dealer's second card:", dhand[1], "(for a total of "+fmt.Sprint(dt)+")")

		if t < 21 {
			for i := 0; dt < 17; i++ { // TODO: Dealer doesn't have high/low ace assessment. Fix that.
				nc, nv := draw()
				dhand = append(dhand, nc)
				dt += nv
				fmt.Println(" - Dealer's next card:", dhand[i+2], "(for a total of "+fmt.Sprint(dt)+")")
			}
		}

		switch {
		case dt == 21 && t < 21:
			fmt.Println("Dealer has blackjack! You Lose. :(")
			showHand(dhand)
		case dt == 21:
			fmt.Println("Dealer has blackjack!")
			showHand(dhand)
		case t > dt && t <= 21:
			fmt.Println("You Win!")
		case t < dt && dt <= 21 && len(hand) < 5:
			fmt.Println("You Lose. :(")
		case dt > 21 && t < 21:
			fmt.Println("Dealer busts.. You Win!! :)")
		case dt == 21 && t < 21:
			fmt.Println("You Lose. :(")
		case t == dt:
			fmt.Println("Push!")
		}

		fmt.Println("Final hand:")
		showHand(hand)

		fmt.Printf("*** Your final score: %d\n\nDealer's final score: %d\n", t, dt)
		fmt.Println("Dealer's final hand:")
		showHand(dhand)
		fmt.Println()

		dd = append(dd, hand...)
		dd = append(dd, dhand...)

		if askYn("Would you like to continue? [Y/n]: ") {
			handcount++
			if handcount == 5 {
				handcount = 0 // Reset handcount.
				fmt.Print("(Re)")
				d = shuffle(append(d, dd...)) // Reshuffle remaining + discarded cards.
				dd = make([]string, 0)        // Empty discarded.
			}
		} else {
			break
		}
	}
}
