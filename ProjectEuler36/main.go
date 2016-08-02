// https://projecteuler.net/problem=36
// The decimal number, 585 = 1001001001 in base 2 (binary), is palindromic in both bases.
// Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.
// (Please note that the palindromic number, in either base, may not include leading zeros.)
// Author: Todd S.
package main

import (
	"fmt"
)

func reverse(r string) string {
	x := []rune(r)
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
	return string(x)
}

func atest(a string) bool {
	return a == reverse(a)
}

func btest(b string) bool {
	return b == reverse(b)
}

func main() {
	var counter int
	for i := 1; i <=1000000; i++ {
		a := fmt.Sprintf("%v", i)
		b := fmt.Sprintf("%b", i)
		if atest(a) == true {
			if btest(b) == true {
				fmt.Printf("%d is palindromic in both base 10 and base 2 (%b).\n", i, i)
				counter += i
			}
		}
	}
	fmt.Println("The total of all numbers between 1 and 1,000,000 "+
	"that are palindromic int both base 2 and base 10 is", counter)
}
