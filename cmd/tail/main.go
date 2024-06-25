package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func main() {
	var nvar int
	flag.IntVar(&nvar, "n", 10, "an int var")
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatal("need file path")
	}
	filepath := flag.Args()[0]
	data, err := os.ReadFile(filepath)
	check(err)
	i := 0
	s := strings.Split(string(data), "\n")
	for j := len(s) - 1; j >= 0 && i < nvar; j-- {
		if s[j] != "" {
			fmt.Println(s[j])
		} else {
			os.Exit(0)
		}
		i++
	}
}
