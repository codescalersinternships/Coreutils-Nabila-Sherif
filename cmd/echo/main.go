package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	flag := false
	args := os.Args[1:]
	if strings.Contains(args[1], "-n") {
		flag = true
		args = args[1:]
	}
	for _, arg := range args {
		fmt.Print(arg, " ")
	}
	if flag {
		fmt.Println()
	}
}
