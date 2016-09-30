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
	bin, e := strconv.ParseInt(b, 2, 0)
	if e != nil {
		fmt.Println("Something went wrong:", e)
		os.Exit(1)
	}
	return int(bin)
}

func intMe(x string) int {
	var i int
	var e error
	if i, e = strconv.Atoi(x); e != nil {
		fmt.Println("Something went wrong:", e)
		os.Exit(1)
	}
	return i
}

func findMe(x, y int) (int, int) {
	count := 8-y
	var i, min, max int
	for i = 256; count > 0; {
		i = i/2; count--
	}
	for j := 0; j <= 256; {
		if x > j {
			j = j+i
		} else {
			max = j-1
			min = j-i
			break
		}
	}
	return min, max
}

func minmax(f, n, t, l int) (string, string) {
	a := fmt.Sprintf("%d.%d.%d.%d", f, n, t, l)
	b := fmt.Sprintf("%08b.%08b.%08b.%08b", f, n, t, l)
	return a, b
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

	var s1, s2, s3, s4 int // Subnet octets
	// First, Next, Third, Last octets + Sub
	f, n, t, l, s := intMe(address[0]), intMe(address[1]), intMe(address[2]), intMe(address[3]), intMe(sub)

	switch s/8 {
	case 0:
		s1 = binMe(s)
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

	var tmin, tmax, min, max int
	var hostmin, hostminb, hostmax, hostmaxb string
	usable := math.Pow(2, float64(32-s))-2
	switch {
	case s >= 24 && s < 31:
		min, max = findMe(l, 32-s)
		hostmin, hostminb = minmax(f, n, t, min+1)
		hostmax, hostmaxb = minmax(f, n, t, max-1)
	case s >= 16 && s < 24:
		min, max = findMe(l, 8)
		tmin, tmax = findMe(t, 24-s)
		hostmin, hostminb = minmax(f, n, tmin, min+1)
		hostmax, hostmaxb = minmax(f, n, tmax, max-1)
	case s < 31:
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

	fmt.Printf("Address:   %d.%d.%d.%d     \t%08b.%08b.%08b.%08b\n", f, n, t, l, f, n, t, l)
	fmt.Printf("Netmask:   %d.%d.%d.%d = %d\t%b.%b.%b.%b\n", s1, s2, s3, s4, s, s1, s2, s3, s4)
	fmt.Printf("HostMin:   %s\t\t%s\n", hostmin, hostminb)
	fmt.Printf("HostMax:   %s\t\t%s\n", hostmax, hostmaxb)
	fmt.Printf("Broadcast: %d.%d.%d.%d     \t%08b.%08b.%08b.%08b\n", f, n, t, max, f, n, t, max)
	fmt.Printf("Hosts/Net: %v\n\n", usable)
}
