// "Winning"* Lotto Number Selector
// Picks 5 numbers between 1 and 69, plus a powerball number
// between 1 and 26.
// *Might not actually pick winning lotto numbers -- use at your own risk.
// Author: Todd S.
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func pbCheck(a []int, b int) bool {
	n := sort.SearchInts(a, b)
	if n < len(a) && a[n] == b {
		return false
	}
	return true
}

func makeRandoms() ([]int,  int, bool) {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := make([]int, 5)
	var z int
	pb := r1.Intn(27-1) + 1 //PowerBall

	for y := 0; y < 5; y++ {
		x[y] = r1.Intn(70-1) + 1
	}

	sort.Ints(x)

	for y2 := 0; y2 < 5; y2++ {
		z += sort.SearchInts(x, x[y2])
		// this should equal exactly 10 if there are no duplicates.
	}

	if z != 10 || pbCheck(x, pb) == false {
		return x, pb, false
	}
	return x, pb, true
}

func main() {
	dupCount := 0
	for z := 0; z < 10; z++ {
		if x, pb, ok := makeRandoms(); ok {
			fmt.Println()
			for y := range x {
				fmt.Printf("%v ", x[y])
			}
			fmt.Printf("\nPowerball number is %d.\n\n----------\n", pb)
			time.Sleep(1000 * time.Millisecond)
		} else {
			dupCount++
		}
	}
	fmt.Printf("\nLines with duplicate numbers (omited): %d\n", dupCount)
}
