package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Card struct {
	face, suit string
	value int
}

func draw() (string, string, int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	c := map[int]string {
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
	x, y, z := draw()
	card1 := Card{x, y, z}
	fmt.Printf("card1: %v\n", card1)
	fmt.Printf("card1.face: %v\n", card1.face)
	fmt.Printf("card1.suit: %v\n", card1.suit)
	fmt.Printf("card1.value: %v\n", card1.value)
}
