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
	arg := os.Args[1]
	if strings.Contains(arg, "-n") {
		arg = os.Args[2]
	}
	data, e := os.ReadFile(arg)
	check(e)
	s := strings.Split(string(data), "\n")
	for j := 0; j < len(s) && j < nvar; j++ {
		if s[j] != "" {
			fmt.Println(s[j])
		} else {
			os.Exit(0)
		}
	}
}
