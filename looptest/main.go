package main

import "fmt"

func scenic() {
	var i int
	i = 1
	for i < 11 {
		fmt.Println("scenic", i)
		i = i+1
	}
}

func succinct() {
	for i := 1 ; i < 11 ; i++ {
		fmt.Println("succinct", i)
	}
}

func main() {
	scenic()
	fmt.Println()
	succinct()
}
