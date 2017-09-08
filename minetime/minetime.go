// minetime sets shutdown and startup times based around avoiding
// PG&E's astronomical peak electricity rates ($0.45/kWh for EV-A).
// Expensive window is only 4 hours on weekends & holidays; 7 hours
// for everything else (14:00-21:00 workdays; 15:00-19:00, otherwise).
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func holidayCheck(date time.Time) bool {
	holidays := []string{
		"2017-09-04",
		"2017-11-11",
		"2017-11-23",
		"2017-12-25",
		"2018-01-01",
		"2018-02-19",
	}
	// (source: https://www.pge.com/tariffs/toudates.shtml)

	for i := range holidays {
		if strings.Contains(fmt.Sprint(date), holidays[i]) {
			return true
		}
	}
	return false
}

func timeSetter(s, file string, wake time.Time, w *bufio.Writer) {
	log.Printf("wakeup is %v (%v)", wake, wake.Unix())
	log.Printf("Setting shutdown to %s.", s)
	cmd := exec.Command("sudo", "shutdown", "-h", s)
	if err := cmd.Run(); err != nil {
		log.Printf("Command finished with error: %v", err)
	}

	_, err := w.WriteString(fmt.Sprint(wake.UTC().Unix()))
	if err != nil {
		log.Fatal(err)
	}
	w.Flush()

	log.Printf("Wrote %v to %s.", wake.UTC().Unix(), file)
}

func main() {
	var wakeup time.Time
	var weekday bool

	file := "/sys/class/rtc/rtc0/wakealarm"

	f, err := os.OpenFile(file,
		os.O_RDWR,
		0644,
	)
	if err != nil {
		log.Fatalf("Unable to open file: %v", file)
	}

	writer := bufio.NewWriter(f)
	_, err = writer.WriteString("")
	if err != nil {
		log.Fatal(err)
	}
	writer.Flush()
	defer f.Close()

	now := time.Now()
	tomorrow := now.Add(7 * time.Hour)

	switch { // weekday/!weekday eval
	case tomorrow.Weekday() == time.Saturday || tomorrow.Weekday() == time.Sunday:
		log.Printf("Tomorrow is %v, so weekday is %t.\n", tomorrow.Weekday(), weekday)
	case holidayCheck(tomorrow):
		log.Printf("Tomorrow is a holiday (%v), so weekday is %t.", fmt.Sprint(tomorrow)[:10], weekday)
	default:
		weekday = true
		log.Printf("Tomorrow is %v, so weekday is %t.\n", tomorrow.Weekday(), weekday)
	}

	if !weekday { // meaning, if weekend or holiday (TODO: send contents to timeSetter()).
		wakeup = time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 19, 00, 00, 00, time.Local)
		timeSetter("15:00", file, wakeup, writer)
	} else {
		wakeup = time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 21, 00, 00, 00, time.Local)
		timeSetter("14:00", file, wakeup, writer)
	}

	byteSlice := make([]byte, 16)
	_, err = f.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	ewake := int64(wakeup.Unix())
	log.Printf("Data read: %v(%v)", string(byteSlice), time.Unix(ewake, int64(0)))
}
