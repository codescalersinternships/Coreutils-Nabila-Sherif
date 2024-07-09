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
	flag.IntVar(&nvar, "n", 10, "for indicating specific number of output lines")
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatal("need file path")
	}
	filepath := flag.Args()[0]
	data, err := os.ReadFile(filepath)
	check(err)
	s := strings.Split(string(data), "\n")
	for i, line := range s {
		if i < nvar {
			fmt.Println(line)
		} else {
			break
		}
	}
}
