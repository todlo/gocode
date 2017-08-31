package main

import (
        "bufio"
        "fmt"
        "log"
        "os"
        "os/exec"
        "time"
)

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
