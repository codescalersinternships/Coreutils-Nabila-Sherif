package main

import (
	"flag"
	"fmt"
)

func main() {
	var showNewLine bool
	flag.BoolVar(&showNewLine, "n", false, "a bool for showing new line")
	flag.Parse()
	for _, arg := range flag.Args() {
		fmt.Print(arg, " ")
	}
	if !showNewLine {
		fmt.Println()
	}
}
