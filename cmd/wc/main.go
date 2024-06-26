package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"unicode/utf8"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func main() {
	args := os.Args[1:]
	var flagL, flagW, flagC bool = false, false, false

	if slices.Contains(args, "-l") {
		flagL = true
	}
	if slices.Contains(args, "-w") {
		flagW = true
	}
	if slices.Contains(args, "-c") {
		flagC = true
	}
	var s []string
	data, err := os.ReadFile(args[len(args)-1])
	check(err)
	s = strings.Split(string(data), "\n")
	l := len(s)
	if flagL {
		fmt.Println("lines : ", l)
	}
	if flagW {
		w := 0
		for j := 0; j < len(s); j++ {
			w += len(strings.Split(s[j], " "))
		}
		fmt.Println("words : ", w)
	}
	if flagC {
		c := utf8.RuneCountInString(string(data))
		fmt.Println("Characters : ", c)
	}
}
