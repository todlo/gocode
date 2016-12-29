package main

import (
	"fmt"
	"regexp"
)

type player struct {
	Name	string
	Round	int
	Positions []int
}

// (getPos)ition on 3x3 board (prompts user for ints 1-9).
func getPos(r int, p string) int {
	var a int
	fmt.Printf("Enter round %d position for player %s: ", r, p)
	if _, e := fmt.Scan(&a); e != nil || a < 1 || a > 9 {
		fmt.Println("Something went wrong: ", e)
		return getPos(r, p)
	}
	return a
}

func posCheck(p int, t []int) bool {
	for i := range t {
		if p == t[i] {
			return true
		}
	}
	return false
}

func board(p int, n string, b *string) {
	re := regexp.MustCompile(fmt.Sprint(p))
	*b = re.ReplaceAllLiteralString(*b, n)
	fmt.Println(*b)
}

func scoreKeep(n string, p int, s [][]int) bool {
	m := map[int][]int{
		1: {0, 0},
		2: {0, 1},
		3: {0, 2},
		4: {1, 0},
		5: {1, 1},
		6: {1, 2},
		7: {2, 0},
		8: {2, 1},
		9: {2, 2},
    }

	if n == "X" {
		s[m[p][0]][m[p][1]] = 1
	} else {
		s[m[p][0]][m[p][1]] = -1
	}

	if winCheck(s) {
		return true
	}
	return false
}

func winCheck(s [][]int) bool {
	var h1, h2, h3, v1, v2, v3, d1, d2 int
	var sums []int

	for i := range s {
		h1 = h1 + s[0][i]
		h2 = h2 + s[1][i]
		h3 = h3 + s[2][i]
		v1 = v1 + s[i][0]
		v2 = v2 + s[i][1]
		v3 = v3 + s[i][2]
		d1 = d1 + s[i][i]
		d2 = d2 + s[i][2-i]
	}

	sums = append(sums, h1, h2, h3, v1, v2, v3, d1, d2)

	for i := range sums {
		if sums[i] == 3 {
			fmt.Println("Win detected for X!")
			return true
		} else if sums [i] == -3 {
			fmt.Println("Win detected for O!")
			return true
		}
	}
	return false
}

func main() {
	var pos int // (pos)itions players want to take.
	var t []int // All positions already (t)aken.
	var win bool // If true, end game.
	b := "1 2 3\n4 5 6\n7 8 9"
	p := make([]player, 2)
	p[0].Name, p[1].Name = "X", "O"

	// s is how wins are determined (see func winCheck).
	s := [][]int{
		[]int{0, 0, 0},
		[]int{0, 0, 0},
		[]int{0, 0, 0},
	}

	fmt.Println(b) // Initial display of empty (b)oard.

	for i := 0; win != true; i++ {
		for j := 0; j < 2; j++ {
			pos = getPos(i+1, p[j].Name)
			p[j].Positions = append(p[j].Positions, pos)
			if posCheck(pos, t) {
				fmt.Println("Position is already taken!", t[:len(t)])
				fmt.Printf("Position(s) already filled by %s: ", p[j].Name)
				fmt.Println(p[j].Positions[:len(p[j].Positions)-1])
				p[j].Positions = p[j].Positions[:len(p[j].Positions)-1]
				if j == 0 { j = 1 } else { j = 0 }
				i -= 1
				continue
			} else {
				t = append(t, pos)
				p[j].Round += 1
				board(pos, p[j].Name, &b)
				win = scoreKeep(p[j].Name, pos, s)
				if win == true {
					fmt.Printf("%s won in %d tries!\n", p[j].Name, p[j].Round)
					break
				}
			}
			if len(t) == 9 {
				fmt.Println("DRAW! Please try again!")
				win = true; break
			}
		}
	}
}
