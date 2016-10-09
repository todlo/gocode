// primeFind is a little program that finds all
// the prime numbers between 2 and 10,000.
// (I made this during a meeting that I really
// should have been paying attention to. Alas.)
// Author: Todd S.
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	var count int
	for i := 2; i <= 10000; i++ {
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
				if count > 2 { break }
			}
		}
		if count <= 2 {
			fmt.Println(i, "is prime.")
		}
		count = 0
	}
	stop := time.Now().UnixNano()
	fmt.Println("Total runtime (in nanoseconds):", stop - start)
}
