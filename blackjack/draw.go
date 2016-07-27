package main

import (
	"fmt"
	"time"
	"math/rand"
	"strings"
)

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
	return fmt.Sprintln(c[num], "of", suit[r.Intn(4)]), num
}

func main() {
	hand := [2]string{}

	card1, score1 := draw()
	if score1 > 10 { score1 = 10 }
	if strings.Contains(card1, "Ace") {
		score1 = score1 + 10
	}
	hand[0] = card1

	card2, score2 := draw()
	if score2 > 10 { score2 = 10 }
	if strings.Contains(card2, "Ace") && score1 < 11 {
		score2 = score2 + 10
	}
	hand[1] = card2

	fmt.Println("Hand:")
	for i := range hand {
		fmt.Print(hand[i])
	}
	fmt.Println("Score:", score1+score2)
}
