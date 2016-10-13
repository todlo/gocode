// vowelsOfHell reverses all the vowels in a word
// given as an argument.
package main

import (
	"fmt"
	"os"
	"strings"
)

//func rev(a []string) []string {
//}

func main() {
	vowels := "aeiouy"
	thearg := os.Args[1]
	bytearg := []rune(thearg)
	fmt.Println(string(bytearg))
	for j := range bytearg {
		for i := 0; i < 6; i++ {
			if strings.IndexRune(string(bytearg[j]), []rune(vowels)[i]) > -1 {
				fmt.Println(string(vowels[i]), "found in position", j)
			}
		}
	}
}
