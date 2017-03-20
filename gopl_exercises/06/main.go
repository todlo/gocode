package main

import (
	"fmt"

	"./intset"
)

func main() {
	var x intset.IntSet
	x.Words = []uint64{1, 5, 1, 6, 9, 1}
	fmt.Println("Preamble:")
	fmt.Printf("x is %v and is of type %T\n", x, x)
	fmt.Println(x.Has(9))
	x.Add(128)
	fmt.Println(x)
	fmt.Println(x.Words)
	fmt.Println(x.Has(64))
	fmt.Println("Exercise 6.1:", x.Len())
	fmt.Println("Exercise 6.2:")
	x.Remove(1)
	fmt.Println(x)
}
