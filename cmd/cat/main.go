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
	if len(os.Args) > 1 {
		var nvar int
		file := os.Args[1]
		printSpecificNum := false
		flag.IntVar(&nvar, "n", 10, "an int var")
		flag.Parse()
		if strings.Contains(os.Args[1], "-n") {
			file = os.Args[2]
			printSpecificNum = true
		}
		data, e := os.ReadFile(file)
		check(e)
		s := strings.Split(string(data), "\n")
		if printSpecificNum {
			for j := 0; j < len(s) && j < nvar; j++ {
				fmt.Println(s[j])
			}
		} else {
			for _, str := range s {
				fmt.Println((str))
			}
		}
	} else {
		os.Exit(1)
	}
}
