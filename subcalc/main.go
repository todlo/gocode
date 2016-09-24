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

func binMe(x int) int {
	b := strings.Replace("00000000", "0", "1", x)
	bin, e := strconv.ParseInt(b, 2, 32)
	if e != nil {
		fmt.Println("Something went wrong:", e)
	}
	return int(bin)
}


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
	var f, n, t, l, s int // First, Next, Third, Last octets + Sub
	var s1, s2, s3, s4 int // Subnet octets
	var e error
	if f, e = strconv.Atoi(address[0]); e != nil {
		fmt.Println("Something went wrong:", e)
		os.Exit(1)
	}
	if n, e = strconv.Atoi(address[1]); e != nil {
		fmt.Println("Something went wrong:", e)
		os.Exit(1)
	}
	if t, e = strconv.Atoi(address[2]); e != nil {
		fmt.Println("Something went wrong:", e)
		os.Exit(1)
	}
	if l, e = strconv.Atoi(address[3]); e != nil {
		fmt.Println("Something went wrong:", e)
		os.Exit(1)
	}
	if s, e = strconv.Atoi(sub); e != nil {
		fmt.Println("Something went wrong:", e)
		os.Exit(1)
	}

	switch s/8 {
	case 1:
		s1 = 255
		if s%8 > 0 { s2 = binMe(s%8) }
	case 2:
		s1, s2 = 255, 255
		if s%8 > 0 { s3 = binMe(s%8) }
	case 3:
		s1, s2, s3 = 255, 255, 255
		if s%8 > 0 { s4 = binMe(s%8) }
	}

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

	fmt.Printf("Address:   %d.%d.%d.%d\t%b.%b.%b.%b\n", f, n, t, l, f, n, t, l)
	fmt.Printf("Netmask:   %d.%d.%d.%d = %d\t%b.%b.%b.%b\n", s1, s2, s3, s4, s, s1, s2, s3, s4)
}
