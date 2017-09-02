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
	"time"
)

func timeSetter(s, file string, wake time.Time, w bufio.ReadWriter) {
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
	tomorrow := now.Add(17 * time.Hour)

	if fmt.Sprint(tomorrow.Weekday()) == "Saturday" || fmt.Sprint(tomorrow.Weekday()) == "Sunday" {
		fmt.Printf("Tomorrow is %v, so weekday is %t.\n", tomorrow.Weekday(), weekday)
	} else {
		weekday = true
		log.Printf("Tomorrow is %v, so weekday is %t.\n", tomorrow.Weekday(), weekday)
	}

	if !weekday { // meaning, if weekend (TODO: or holiday).
		wakeup = time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 19, 00, 00, 00, time.Local)
		log.Printf("wakeup is %v (%v)\n", wakeup, wakeup.Unix())
		log.Print("Setting shutdown to 15:00.")
		cmd := exec.Command("sudo", "shutdown", "-h", "15:00")
		if err := cmd.Run(); err != nil {
			log.Printf("Command finished with error: %v", err)
		}

		_, err := writer.WriteString(fmt.Sprint(wakeup.UTC().Unix()))
		if err != nil {
			log.Fatal(err)
		}
		writer.Flush()

		log.Printf("Wrote %v to %s.", wakeup.UTC().Unix(), file)

	} else {
		wakeup = time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 21, 00, 00, 00, time.Local)
		log.Printf("wakeup is %v (%v)\n", wakeup, wakeup.Unix())
		log.Print("Setting shutdown to 14:00.")
		cmd := exec.Command("sudo", "shutdown", "-h", "14:00")
		if err := cmd.Run(); err != nil {
			log.Printf("Command finished with error: %v", err)
		}

		_, err := writer.WriteString(fmt.Sprint(wakeup.UTC().Unix()))
		if err != nil {
			log.Fatal(err)
		}
		writer.Flush()

		log.Printf("Wrote %v to %s.", wakeup.UTC().Unix(), file)
	}

	byteSlice := make([]byte, 16)
	_, err = f.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	ewake := int64(wakeup.Unix())
	log.Printf("Data read: %v(%v)", string(byteSlice), time.Unix(ewake, int64(0)))
}
