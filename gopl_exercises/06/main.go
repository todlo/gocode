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
	fmt.Println(x.Has(64))
	fmt.Println("Exercise 6.1.1:", x.Len())
	fmt.Println("Exercise 6.1.2:")
	x.Remove(1)
	fmt.Println(x)
	fmt.Println("Exercise 6.1.3:")
	x.Clear()
	fmt.Println(x)
	x.Words = []uint64{1, 5, 1, 6, 9, 1}
	fmt.Println("Exercise 6.1.4:")
	fmt.Println(&x.Words)
	x.Copy()
	fmt.Println(x)
}
