// subcalc takes a given IPv4 address with subnet and returns
// the range of usable addresses in that subnet.

package main

import (
	"fmt"
	"math"
	"os"
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
	var f, n, t, l, s int // First, Next, Third, Last octets + sub
	var e error
	if f, e = strconv.Atoi(address[0]); e != nil {
		fmt.Println("Something went wrong:", e)
		os.Exit(1)
	}
	n, _ = strconv.Atoi(address[1])
	t, _ = strconv.Atoi(address[2])
	l, _ = strconv.Atoi(address[3])
	s, _ = strconv.Atoi(sub)

	switch {
	case s < 31:
		usable := math.Pow(2, float64(32-s))-2
		fmt.Printf("There are %v usable addresses in a /%d subnet.\n", usable, s)
	case s == 31:
		var o int // other end
		if l%2 == 0 {
			o = l+1
		} else {
			o = l-1
		}
		fmt.Printf("This is a point-to-point link (RFC 3021), "+
			"the other end being %d.%d.%d.%d/%d.\n", f, n, t, o, s)
	default:
		fmt.Println("/32 (255.255.255.255) is a device address; nothing to calculate!")
	}
}
