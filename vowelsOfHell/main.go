// vowelsOfHell reverses all the vowels in a word
// given as an argument.
package main

import (
	"fmt"
	"os"
	"strings"
)

func rev(a []rune) []rune {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func main() {
	vowels := "aeiouy"
	thearg := []byte(os.Args[1])
	var vcount int
	var argv []rune
	var index []int
	for j := range []rune(os.Args[1]) {
		for i := 0; i < 6; i++ {
			if strings.IndexRune(string(os.Args[1][j]), []rune(vowels)[i]) > -1 {
				fmt.Println(string(vowels[i]), "found in position", j)
				vcount++
				argv = append(argv, rune(vowels[i]))
				index = append(index, j)
			}
		}
	}
	fmt.Printf("There are %d vowels in the word \"%s\".\n", vcount, os.Args[1])
	rargv := rev(argv)
	for i := range rargv {
		thearg[index[i]] = byte(rargv[i])
	}
	fmt.Println(string(thearg))
}
