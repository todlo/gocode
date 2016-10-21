// wordcounter counts how many times a word occurs in a given text. #homework
// Author: Todd S.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/tmp/test_text.txt")
	if err != nil {
		fmt.Println("Error opening file:", f)
		os.Exit(1)
	}
	defer f.Close()

	var words []string
	wordmap := make(map[string]int)

	scanMe := bufio.NewScanner(f)
	scanMe.Split(bufio.ScanWords)

	for scanMe.Scan() {
		words = append(words, scanMe.Text())
	}

	for i := range words {
		_, v := wordmap[words[i]]
		if v {
			wordmap[words[i]] = wordmap[words[i]]+1
			fmt.Println("DEBUG: found key, value:", words[i], wordmap[words[i]])
		} else {
			wordmap[words[i]] = 1
			fmt.Println("DEBUG: new key detected. adding key, value:", words[i], wordmap[words[i]])
		}
	}
	for k, v := range wordmap {
		fmt.Println(k, v)
	}
}
