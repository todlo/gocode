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

func main() {
	var roll1, roll2, score int
	frame := 1
	p1scores := []int{5, 3, 6, 4, 5, 4, 7, 1, 5, 5, 9, 0, 3, 7, 10, 8, 0, 7, 2}
	fmt.Println("p1scores to process:", p1scores)
	for frame < 10 {
		roll1, roll2 = p1scores[0], p1scores[1]
		switch {
		case roll1 + roll2 < 10:
			score += roll1 + roll2
			p1scores = p1scores[2:]
		case roll1 + roll2 == 10:
			score += roll1 + roll2 + p1scores[2]
			p1scores = p1scores[2:]
		case roll1 == 10:
			score += 10 + p1scores[1] + p1scores[2]
			p1scores = p1scores[1:]
		}
		fmt.Printf("p1scores left: %d\nFrame: %d\nScore: %d\n", p1scores, frame, score)
		frame++
	}
	fmt.Println("Frame: 10")
	for i := range p1scores {
		score += p1scores[i]
	}
	fmt.Println("Final Score:", score)
}
