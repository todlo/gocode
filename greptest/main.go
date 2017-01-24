package main

import (
	"fmt"
	"os/exec"
	"regexp"
)

func main() {
	cmd, err := exec.Command("ls").Output()
	if err != nil {
		fmt.Printf("Something went wrong: %v\n", err)
	}

	re := regexp.MustCompile("bb.*")
	fmt.Println("Looking for", re)

	bb := re.FindAllString(string(cmd), -1)

	for i := range bb {
		fmt.Println(bb[i])
	}
}
