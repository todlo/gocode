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

type player struct {
	name string
}

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
		case roll1 < 10 && roll1 + roll2 == 10:
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

func getPlayers() int {
	var a int
	fmt.Print("How many players? [1-8]: ")
	if _, err := fmt.Scan(&a); err != nil || a < 1 || a > 8 {
		fmt.Println(fmt.Sprint(err)+": Please pick a number from 1 through 8.")
		getPlayers()
	}
	return a
}

func getNames(p int) string {
	var a string
	fmt.Printf("What is the name of player %d? ", p)
	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println(err, "Something went wrong.")
		getNames(p)
	}
	return a
}

func getFrames(name string) []int {
	set := make([]int, 21)
	var index int
	// Frames 1-9:
	for i := 0; i < 9; i++ {
		fmt.Println("Frame", i+1)
		for j := 1; j <= 2; j++ {
			fmt.Printf("Please enter %s's score for frame %d, ball %d [default 0]: ", name, i+1, j)
			_, err := fmt.Scanln(&set[index])
			switch {
			case err != nil && fmt.Sprint(err) != "unexpected newline":
				fmt.Println(err, "Something went wrong.")
				j = 1; i--
			case set[index] == 10 && j == 1:
				j++
				index += 2
				break
			case fmt.Sprint(err) == "unexpected newline":
				fmt.Printf("Setting index %d to 0.\n", set[index])
				set[index] = 0
				index++
			default:
				index++
			}
		}
	}
	// Frame 10:
	fmt.Println("Frame 10")
	index = 18
	for j := 1; j <= 3; j++ {
		fmt.Printf("Please enter %s's score for frame 10, ball %d [default 0]: ", name, j)
		if _, err := fmt.Scan(&set[index]); err != nil && fmt.Sprint(err) != "unexpected newline" {
			fmt.Println(err, "Something went wrong.")
			j = 1
			index = 18
		} else if j == 2 && set[18] + set[19] < 10 {
			j = 3
			break
		} else {
			index++
		}
	}
	return set
}


func main() {
	numplayers := getPlayers()
	x := make([]string, numplayers)
	fmt.Printf("You entered %d players.\n", numplayers)
	for i := 0; i < numplayers; i++ {
		x[i] = fmt.Sprintf(getNames(i+1))
		fmt.Printf("Player %d is %s.\n", i+1, x[i])
		frames := getFrames(x[i])
		fmt.Printf("%s's set: %v\n", x[i], frames)
		fmt.Printf("%s's total:\n", x[i])
		bowl(frames)
	}
	//p1 := player{fmt.Sprint(x[0])}
	//fmt.Println(p1.name)
	//p1scores := []int{5, 3, 6, 4, 5, 4, 7, 1, 5, 5, 9, 0, 3, 7, 10, 8, 0, 7, 2}
	//p2scores := []int{5, 3, 10, 5, 4, 7, 1, 5, 5, 9, 0, 3, 7, 10, 8, 0, 7, 2}
	//p3scores := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	//bowl(p1scores)
	//bowl(p2scores)
	//bowl(p3scores)
}
