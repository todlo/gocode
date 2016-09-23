// subcalc takes a given IPv4 address with subnet and returns
// the range of usable addresses in that subnet.

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	address := make([]string, 4)
	var a string
	fmt.Print("Enter v4 address with subnet (e.g., 1.2.3.4/24): ")
	fmt.Scanln(&a)
	sub := a[strings.Index(a, "/")+1:]
	a = a[:strings.Index(a, "/")]
	for i := 0; i < 3; i++ {
		address[i] = a[:strings.Index(a, ".")]
		a = a[strings.Index(a, ".")+1:]
	}
	address[3] = a

	s, _ := strconv.Atoi(sub)
	switch {
	case s < 31:
		usable := math.Pow(2, float64(32-s))-2
		fmt.Printf("There are %v usable addresses in a /%d subnet.\n", usable, s)
	case s == 31:
		fmt.Printf("This is a point-to-point address, the other end being %s.\n")
	default:
		fmt.Println("/32 (255.255.255.255) is a device address; nothing to calculate!")
	}

	fmt.Println(address)
	for x := range address {
		i, _ := strconv.Atoi(address[x])
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
