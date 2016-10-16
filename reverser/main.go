package main

import (
	"fmt"
	"os"
)

func revMe(a []rune) string {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}

func main() {
	thearg := []rune(os.Args[1])
	fmt.Println(revMe(thearg))
}
