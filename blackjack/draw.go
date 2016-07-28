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

func draw() (string, int) {
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
	return c[num] + " of " + suit[r.Intn(4)], num
}

func main() {
	var highace bool
	hand := make([]string, 2)

	card1, score1 := draw()
	if score1 > 10 { score1 = 10 }
	if strings.Contains(card1, "Ace") {
		score1 = score1 + 10
		highace = true
	}
	hand[0] = card1

	card2, score2 := draw()
	if score2 > 10 { score2 = 10 }
	if strings.Contains(card2, "Ace") && score1 < 11 {
		score2 = score2 + 10
		highace = true
	}
	hand[1] = card2

	t := score1+score2
	fmt.Println("*** Hand: ***")
	for i := range hand {
		fmt.Println(hand[i])
	}
	fmt.Println("Score:", t)

	for t < 21 {
		if askYn("Would you like to hit? [Y/n]: ") {
			newcard, value := draw()
			if value > 10 { value = 10 }
			switch {
				case strings.Contains(newcard, "Ace") && t + 11 <= 21: 
					highace = true
					t += 11
				case highace == true && t > 21: 
					highace = false
					t -= 10
				default:
					t += value
			}
			hand = append(hand, newcard)
			fmt.Printf("Dealer deals a %s.\nNew score: %d\n", newcard, t)
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
