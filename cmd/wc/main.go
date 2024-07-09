package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func main() {
	var flagL bool
	flag.BoolVar(&flagL, "l", false, "a bool for output lines")
	var flagW bool
	flag.BoolVar(&flagW, "w", false, "a bool for output lines")
	var flagC bool
	flag.BoolVar(&flagC, "c", false, "a bool for output lines")
	flag.Parse()
	filepath := flag.Args()[0]
	var s []string
	data, err := os.ReadFile(filepath)
	check(err)
	s = strings.Split(string(data), "\n")
	l := len(s)
	if flagL || (!flagL && !flagC && !flagW) {
		fmt.Print(l, " ")
	}
	if flagW || (!flagL && !flagC && !flagW) {
		w := 0
		for j := 0; j < len(s); j++ {
			w += len(strings.Split(s[j], " "))
		}
		fmt.Print(w, " ")
	}
	if flagC || (!flagL && !flagC && !flagW) {
		c := utf8.RuneCountInString(string(data))
		fmt.Println(c, " ")
	}
	fmt.Print(filepath)
}
