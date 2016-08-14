// blackjack draws from a single- stack deck.// TODO: Add option to pull from single OR 9-stack deck.
// Deals initial 2 cards, then asks user if they'd like to hit (default is yes) until
// user wins (score == 21; blackjack), busts (score > 21; bust), gets 5 cards, or chooses
// to stay.
// Author: Todd S.
package main

import (
	"fmt"
	"strings"

	"./deck"
)

var d = deck.Deck()

type Card struct {
	card string
	value int
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

func draw() (string, int) {
	topcard := d[0]
	d = d[1:]
	value := cardEval(topcard)
	return topcard, value
}

func handInit() ([]string, int, bool) {
	var highace bool
	hand := make([]string, 2)

	x, y := draw()
	c1 := Card{x, y}
	x, y = draw()
	c2 := Card{x, y}

	if strings.Contains("Ace", c1.card) {c1.value = 11 ; highace = true }
	if strings.Contains("Ace", c2.card) && c1.value <11 { c2.value = 11 ; highace = true }

	hand[0] = c1.card
	hand[1] = c2.card

	t := c1.value + c2.value

	return hand, t, highace
}

func main() {
	hand, t, highace := handInit()

	fmt.Println("*** Your Hand: ***")
	for i := range hand {
		fmt.Printf("%d. %s\n", i+1, hand[i])
	}

	fmt.Println("*** Dealer's Hand: ***")
	dhand, dt, _ := handInit()
	if dt == 21 {
		fmt.Printf("Dealer has blackjack!\n1. %s\n2. %s\n", dhand[0], dhand[1])
	} else {
		fmt.Printf("1. %s\n2. (Face Down)\n", dhand[0])
	}

	fmt.Println("Score:", t)
	if t == 21 { fmt.Println("BLACKJACK!! :D") }

	for t < 21 {
		if askYn("Would you like to hit? [Y/n]: ") {
			nc, nv := draw()
			switch {
				case strings.Contains("Ace",nc) && t + 11 <= 21:
					highace = true
					t += 11
				default:
					t += nv
			}
			hand = append(hand, nc)
			fmt.Printf("*** Your next card: %s.\n", nc)
		} else {
			fmt.Println(" - Dealer's second card:", dhand[1], "( for a total of", dt, ")")
			break
		}
		if highace == true && t > 21 {
			fmt.Println("High Ace becomes Low Ace...")
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
			fmt.Println("BUSTED! :(")
		default:
			fmt.Println("Current score:", t)
		}
	}
	if t < 21 {
		for i := 0; dt < 17; i++ {
			nc, nv := draw()
			dhand = append(dhand, nc)
			dt += nv
			fmt.Println(" - Dealer's next card:", dhand[i+2], "( for a total of", dt, ")")
		}
	}
	switch {
	case t > dt && t <= 21 :
		fmt.Println("You Win!")
	case t < dt && dt <= 21 && len(hand) < 5:
		fmt.Println("You Lose. :(")
	case dt > 21 && t < 21:
		fmt.Println("Dealer busts.. You Win!! :)")
	case t == dt:
		fmt.Println("Push!")
	}
	fmt.Println("Final hand:")
	for i := range hand {
		fmt.Printf("%d. %s\n", i+1, hand[i])
	}
	fmt.Println("Final score:", t)
	fmt.Print("Dealer's score: ", dt, " ( ")
	for i := range dhand {
		fmt.Printf("%s ", dhand[i])
	}
	fmt.Printf(")\n")
	fmt.Println("DEBUG: len(deck) is", len(d))
}
