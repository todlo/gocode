// blackjack draws from an "infinite" stack // TODO: Add option to pull from single or 9-stack deck.
// Deals initial 2 cards, then asks user if they'd like to hit (default is yes) until
// user wins (score == 21; blackjack), busts (score > 21; bust), gets 5 cards, or chooses
// to stay.
// Author: Todd S.
package main

import (
	"fmt"
	"time"
	"math/rand"
	"strings"
)

type Card struct {
	face, suit string
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

func draw() (string, string, int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
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
	num := r.Intn(14-1) + 1
	suit := [4]string{"Spades", "Diamonds", "Clubs", "Hearts"}
	return c[num], suit[r.Intn(4)], num
}

func main() {
	var highace bool
	hand := make([]string, 2)

	x, y, z := draw()
	c1 := Card{x, y, z}
	x, y, z = draw()
	c2 := Card{x, y, z}

	if c1.value > 10 { c1.value = 10 }
	if c2.value > 10 { c2.value = 10 }
	if c1.face == "Ace" { c1.value = 11 ; highace = true }
	if c2.face == "Ace" && c1.value <11 { c2.value = 11 ; highace = true }

	hand[0] = fmt.Sprint(c1.face, " of ", c1.suit)
	hand[1] = fmt.Sprint(c2.face, " of ", c2.suit)

	t := c1.value + c2.value

	fmt.Println("*** Hand: ***")
	for i := range hand {
		fmt.Println(hand[i])
	}
	fmt.Println("Score:", t)

	for t < 21 {
		if askYn("Would you like to hit? [Y/n]: ") {
			nf, ns, nv := draw()
			if nv > 10 { nv = 10 }
			switch {
				case nf == "Ace" && t + 11 <= 21:
					highace = true
					t += 11
				case highace == true && t > 21:
					highace = false
					t -= 10
				default:
					t += nv
			}
			hand = append(hand, fmt.Sprint(nf, " of ", ns))
			fmt.Printf("You are dealt a %s of %s.\nNew score: %d\n", nf, ns, t)
		} else {
			break
		}
		switch {
		case highace == true && t > 21:
			fmt.Println("High Ace becomes Low Ace...")
			highace = false
			t -= 10
			fmt.Println("New score:", t)
		case t == 21:
			fmt.Println("BLACKJACK!! :D")
		case t > 21:
			fmt.Println("BUSTED! :(")
		case t < 21 && len(hand) == 5:
			fmt.Println("5-card hand... YOU WIN!! :D")
		}
	}
	fmt.Println("Final hand:")
	for i := range hand {
		fmt.Println(hand[i])
	}
	fmt.Println("Final score:", t)
}
