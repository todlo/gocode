// minetime sets shutdown and startup times based around avoiding
// PG&E's astronomical peak electricity rates ($0.45/kWh for EV-A).
// Expensive window is only 4 hours on weekends & holidays; 7 hours
// for everything else (14:00-21:00 workdays; 15:00-19:00, otherwise).
package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

type h struct {
	hol  time.Time
	name string
}

func holidayCheck(date time.Time) bool {
	holidays := []string{
		"2018-01-01", // New Years Day
		"2018-02-19", // Presidents Day
		"2018-05-28", // Memorial Day
		"2018-07-04", // Independence Day
		"2018-09-03", // Labor Day
		"2018-11-12", // Veterans Day
		"2018-11-22", // Thanksgiving
		"2018-12-25", // Christmas
	}
	// (source: https://www.pge.com/tariffs/toudates.shtml)

	for i := range holidays {
		if strings.Contains(fmt.Sprint(date), holidays[i]) {
			return true
		}
	}
	return false
}

func timeSetter(s string, wake time.Time) {
	log.Printf("wakeup is %v (%v)", wake, wake.Unix())
	log.Printf("Setting shutdown to %s.", s)
	cmd := exec.Command("sudo", "shutdown", "-h", s)
	if err := cmd.Run(); err != nil {
		log.Printf("Command finished with error: %v", err)
	}

	w := fmt.Sprint(wake.UTC().Unix())
	cmd = exec.Command("sudo", "rtcwake", "-m", "no", "-t", w, "-u")
	if err := cmd.Run(); err != nil {
		log.Printf("Command finished with error: %v", err)
	}

	log.Printf("Wrote %v to /dev/rtc0.", wake.UTC().Unix())
}

func main() {
	var wakeup time.Time
	var weekday bool

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

	if !weekday { // meaning, if weekend or holiday
		wakeup = time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 19, 00, 00, 00, time.Local)
		timeSetter("15:00", wakeup)
	} else {
		wakeup = time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 21, 00, 00, 00, time.Local)
		timeSetter("14:00", wakeup)
	}
}
