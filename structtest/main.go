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
	var highace bool

	x, y, z := draw()
	c1 := Card{x, y, z}
	x, y, z = draw()
	c2 := Card{x, y, z}

	if c1.value > 10 { c1.value = 10 }
	if c2.value > 10 { c2.value = 10 }
	if c1.face == "Ace" { c1.value = 11 ; highace = true }
	if c2.face == "Ace" && c1.value <11 { c2.value = 11 ; highace = true }

	fmt.Printf("%s of %s\n", c1.face, c1.suit)
	fmt.Printf("c1.face: %v\n", c1.face)
	fmt.Printf("c1.suit: %v\n", c1.suit)
	fmt.Printf("c1.value: %v\n", c1.value)

	fmt.Printf("%s of %s\n", c2.face, c2.suit)
	fmt.Printf("c2.face: %v\n", c2.face)
	fmt.Printf("c2.suit: %v\n", c2.suit)
	fmt.Printf("c2.value: %v\n", c2.value)

	fmt.Printf("Total: %d\n", c1.value + c2.value)
	fmt.Printf("Ace detected? %v\n", highace)
}
