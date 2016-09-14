// bowling.go adds the player scores from a standard game of US-style
// bowling (10 frames, 2 rolls per frame to try to knock down 10 pins).
// Right now it just takes a hardcoded array of single player scores
// (p1scores == "player one scores"), but the idea is to input scores
// from as many as 8 players and have them tallied (and then ranked by
// player).
package main

import (
	"fmt"
)

func bowl(set []int) {
	var roll1, roll2, score int
	frame := 1
	fmt.Println("scores to process:", set)
	for frame < 10 {
		fmt.Printf("Frame: %d\n", frame)
		roll1, roll2 = set[0], set[1]
		switch {
		case roll1 + roll2 < 10:
			score += roll1 + roll2
			set = set[2:]
		case roll1 + roll2 == 10:
			fmt.Println("SPARE!")
			score += roll1 + roll2 + set[2]
			set = set[2:]
		case roll1 == 10:
			fmt.Println("S T R I K E !!")
			score += 10 + set[1] + set[2]
			set = set[1:]
		}
		fmt.Printf("Score: %d\n\n", score)
		frame++
	}
	fmt.Println("Frame: 10")
	for i := range set {
		score += set[i]
	}
	fmt.Println("Final Score:", score)
}

func main() {
	p1scores := []int{5, 3, 6, 4, 5, 4, 7, 1, 5, 5, 9, 0, 3, 7, 10, 8, 0, 7, 2}
	p2scores := []int{5, 3, 10, 5, 4, 7, 1, 5, 5, 9, 0, 3, 7, 10, 8, 0, 7, 2}
	p3scores := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	bowl(p1scores)
	bowl(p2scores)
	bowl(p3scores)
}
