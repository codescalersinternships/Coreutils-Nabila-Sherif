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
		var outputlines bool
		flag.BoolVar(&outputlines, "n", false, "a bool for output lines")
		flag.Parse()
		filepath := flag.Args()[0]
		data, e := os.ReadFile(filepath)
		check(e)
		s := strings.Split(string(data), "\n")
		for i, str := range s {
			if outputlines {
				fmt.Println(i, " ", str)
			} else {
				fmt.Println((str))
			}
		}
	} else {
		os.Exit(1)
	}
}
