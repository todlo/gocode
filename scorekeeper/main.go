package main

import (
	"fmt"
)

type Player struct {
	Name string
	Total int
}

func getPlayers(a *int) int {
	fmt.Print("How many players? [1-8]: ")
	if _, e := fmt.Scan(a); e != nil || *a < 1 || *a > 8 {
		fmt.Println(fmt.Sprint(e)+": Please pick a number from 1 through 8.")
		getPlayers(a)
	}
	return *a
}

func getRounds(r *int) int {
	fmt.Print("How many rounds? ")
	if _, e := fmt.Scan(r); e != nil || *r < 1 {
		fmt.Println(fmt.Sprint(e)+": Please pick a positive while number.")
		getRounds(r)
	}
	return *r
}

func getNames(p int) string {
	var a string
	fmt.Printf("What is the name of player %d? ", p)
	if _, e := fmt.Scanln(&a); e != nil {
		fmt.Println("Something went wrong:", e)
		getNames(p)
	}
	return a
}

func getScore(r int, p string) int {
	var a int
	fmt.Printf("Enter round %d score for player %s: ", r, p)
	if _, e := fmt.Scan(&a); e != nil {
		fmt.Println("Something went wrong: ", e)
		getScore(r, p)
	}
	return a
}

func main() {
	var n, r, s int // players, rounds, scores
	n = getPlayers(&n)
	players := make([]Player, n) // Creates a slice of type main.Player
	r = getRounds(&r)
	for i := 0; i < n; i++ {
		players[i].Name = getNames(i+1)
	}

	for f := 0; f < r; f++ { // f = frame
		for i := range players {
			s = getScore(f+1, players[i].Name)
			players[i].Total += s
		}
	}

	for i := range players {
		fmt.Printf("%s's total: %d\n", players[i].Name, players[i].Total)
	}
}
