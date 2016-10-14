// vowelsOfHell reverses all the vowels in a word
// given as an argument.
package main

import (
	"fmt"
	"os"
	"strings"
)

// Rev takes a slice of runes, reverses the order,
// then returns that reversed slice.
func rev(a []rune) []rune {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func main() {
	vowels := "aeiouy"
	thearg := []rune(os.Args[1]) // The word we're working on in the form of a slice of runes.
	var vcount int // Vowel count of the word we're working on.
	var argv []rune // Vowels from the word in the form of a slice of runes.
	var index []int // Index so that we know where to reinsert reversed []runes.
	for j := range thearg {
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
		thearg[index[i]] = rargv[i]
	}
	fmt.Println(string(thearg))
}
