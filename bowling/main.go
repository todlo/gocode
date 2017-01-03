// bowling.go adds the player scores from a standard game of US-style
// bowling (10 frames, 2 rolls per frame to try to knock down 10 pins).
// Takes input for up to 8 players, with the user entering the full set 
// of each player. Player set goes into a []int (slice of ints). Scoring
// algorithm lives in bowl(), which is called with that []int set. 
package main

import (
	"fmt"
)

type player struct {
	Name string
	Frame []int
}

func bowl(set []int) {
	var roll1, roll2, score int
	frame := 1
	fmt.Println("frames to process:", set)
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
		case roll1 == 10 && set[2] < 10:
			fmt.Println("S T R I K E !!")
			score += 10 + set[2] + set[3]
			set = set[2:]
		case roll1 == 10 && set[2] == 10:
			fmt.Println("S T R I K E !!")
			score += 10 + set[2] + set[4]
			set = set[2:]
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

func getPlayers(a *int) int {
	fmt.Print("How many players? [1-8]: ")
	if _, e := fmt.Scan(a); e != nil || *a < 1 || *a > 8 {
		fmt.Println(fmt.Sprint(e)+": Please pick a number from 1 through 8.")
		getPlayers(a)
	}
	return *a
}

func getNames(p int) string {
	var a string
	fmt.Printf("What is the name of player %d? ", p)
	if _, e := fmt.Scanln(&a); e != nil {
		fmt.Println(e, "Something went wrong.")
		getNames(p)
	}
	return a
}

func getFrames(frame int, name string) []int {
	pins := make([]int, 2)
	// Frames 1-9:
	if frame <= 9 {
		for j := 0; j < 2; j++ {
			fmt.Printf("~~ Frame %d.%d ~~\n", frame, j+1)
			fmt.Printf("Please enter %s's score (0-10) for frame %d, ball %d [default 0]: ", name, frame, j+1)
			_, e := fmt.Scanln(&pins[j])
			switch {
			case fmt.Sprint(e) == "unexpected newline":
				pins[j] = 0
			case e != nil && fmt.Sprint(e) != "unexpected newline":
				fmt.Println(e, "...backing up.")
				if j == 1 {
					j = 0
				} else {
					j = 1
				}
				continue
			case j == 0 && pins[j] > 10:
				fmt.Println("There are only 10 pins! Please enter a number from 0 through 10")
				j = 0
				continue
			case j == 1 && pins[j] > 10-pins[0]:
				fmt.Printf("There are only %d pins left! Please enter a number from 0 through %d.\n", 10-pins[0], 10-pins[0])
				j = 0
				continue
			case pins[j] == 10 && j == 0:
				j++
				break
			default:
				continue
			}
		}
		return pins
	} else {
	// Frame 10:
	fmt.Println("Frame 10")
	pins = append(pins, 0)
		for j := 0; j < 3; j++ {
			fmt.Printf("Please enter %s's score for frame 10, ball %d [default 0]: ", name, j+1)
			if _, err := fmt.Scanln(&pins[j]); err != nil && fmt.Sprint(err) != "unexpected newline" {
				fmt.Println(err, "Something went wrong.")
				j = 0
			} else if j == 1 && pins[0] + pins[1] < 10 {
				j = 3
				break
			} else {
				continue
			}
		}
	}
	return pins
}


func main() {
	var a int
	n := getPlayers(&a)
	p := make([]player, n)
	fmt.Printf("You entered %d players.\n", n)
	for i := 0; i < n; i++ {
		p[i].Name = getNames(i+1)
		fmt.Printf("Player %d is %s.\n", i+1, p[i].Name)
	}
	for f := 1; f <= 10; f++ {
		for i := range p {
			x := getFrames(f, p[i].Name)
			p[i].Frame = append(p[i].Frame, x...)
		}
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%s's full set: %v\n", p[i].Name, p[i].Frame)
		bowl(p[i].Frame)
	}
}
