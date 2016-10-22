// subcalc takes a given IPv4 address with subnet and returns
// the range of usable addresses in that subnet.

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

func getAddr(a string) ([]string, string) {
	address := make([]string, 4)
	if a == "" {
		fmt.Print("Enter v4 address with subnet (e.g., 1.2.3.4/24): ")
		fmt.Scanln(&a)
	}
	sub := a[strings.Index(a, "/")+1:]
	a = a[:strings.Index(a, "/")]
	for i := 0; i < 3; i++ {
		address[i] = a[:strings.Index(a, ".")]
		a = a[strings.Index(a, ".")+1:]
	}
	address[3] = a
	return address, sub
}

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
		if x >= j {
			j = j+i
		} else {
			max = j-1
			min = j-i
			break
		}
	}
	return min, max
}

func recomp(f, n, t, l int) (string, string) {
	a := fmt.Sprintf("%d.%d.%d.%d", f, n, t, l)
	b := fmt.Sprintf("%08b.%08b.%08b.%08b", f, n, t, l)
	return a, b
}

func main() {
	var address []string
	var sub string
	if len(os.Args) == 2 {
		address, sub = getAddr(os.Args[1])
	} else {
		address, sub = getAddr("")
	}

	var s1, s2, s3, s4 int // Subnet octets
	// First, Next, Third, Last octets + Sub
	f, n, t, l, s := intMe(address[0]), intMe(address[1]), intMe(address[2]), intMe(address[3]), intMe(sub)

	var nmin, nmax, tmin, tmax, min, max int
	var network, networkb, hostmin, hostminb, hostmax, hostmaxb, bcast, bcastb string
	usable := math.Pow(2, float64(32-s))-2

	switch s/8 {
	case 0:
		s1 = binMe(s)
	case 1:
		s1 = 255
		if s%8 > 0 { s2 = binMe(s%8) }
		min, max = findMe(l, 8)
		tmin, tmax = findMe(t, 8)
		nmin, nmax = findMe(n, 16-s)
		hostmin, hostminb = recomp(f, nmin, tmin, min+1)
		hostmax, hostmaxb = recomp(f, nmax, tmax, max-1)
		network, networkb = recomp(f, nmin, tmin, min)
		bcast, bcastb = recomp(f, nmax, tmax, max)
	case 2:
		s1, s2 = 255, 255
		if s%8 > 0 { s3 = binMe(s%8) }
		min, max = findMe(l, 8)
		tmin, tmax = findMe(t, 24-s)
		hostmin, hostminb = recomp(f, n, tmin, min+1)
		hostmax, hostmaxb = recomp(f, n, tmax, max-1)
		network, networkb = recomp(f, n, tmin, min)
		bcast, bcastb = recomp(f, n, tmax, max)
	case 3:
		s1, s2, s3 = 255, 255, 255
		if s%8 > 0 { s4 = binMe(s%8) }
		min, max = findMe(l, 32-s)
		if s == 31 {
			hostmin, hostminb = recomp(f, n, t, min)
			hostmax, hostmaxb = recomp(f, n, t, max)
			network, networkb = hostmin, hostminb
			bcast, bcastb = hostmax, hostmaxb
		} else {
			hostmin, hostminb = recomp(f, n, t, min+1)
			hostmax, hostmaxb = recomp(f, n, t, max-1)
			network, networkb = recomp(f, n, t, min)
			bcast, bcastb = recomp(f, n, t, max)
		}
	case 4:
		fmt.Println("/32 (255.255.255.255) is a device address; nothing to calculate!")
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 8, ' ', 0)
	fmt.Fprintf(w, "Address:   %d.%d.%d.%d\t%08b.%08b.%08b.%08b\n", f, n, t, l, f, n, t, l)
	fmt.Fprintf(w, "Netmask:   %d.%d.%d.%d = %d\t%b.%b.%b.%b\n", s1, s2, s3, s4, s, s1, s2, s3, s4)
	fmt.Fprintf(w, "Network:   %s/%d\t%s\n", network, s, networkb)
	fmt.Fprintf(w, "HostMin:   %s\t%s\n", hostmin, hostminb)
	fmt.Fprintf(w, "HostMax:   %s\t%s\n", hostmax, hostmaxb)
	fmt.Fprintf(w, "Broadcast: %s\t%s\n", bcast, bcastb)
	fmt.Fprintf(w, "Hosts/Net: %v\n\n", int(usable))
	w.Flush()
}
