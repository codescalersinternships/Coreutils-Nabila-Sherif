package main

import (
	"fmt"
	"os"
)

func main() {
	showNewLine := true
	args := os.Args[1:]
	if args[0] == "-n" {
		showNewLine = false
		args = args[1:]
	}
	for _, arg := range args {
		fmt.Print(arg, " ")
	}
	if showNewLine {
		fmt.Println()
	}
}
