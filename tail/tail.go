package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	nPtr := flag.Int("n", 10, "an int")
	flag.Parse()
	arg := os.Args[1]
	if strings.Contains(arg, "-") {
		arg = os.Args[2]
	}
	data, _ := os.ReadFile(arg)
	var s []string
	i := 0
	s = strings.Split(string(data), "\n")
	for j := len(s) - 1; j >= 0 && i < *nPtr; j-- {
		if s[j] != "" {
			fmt.Println(s[j])
		} else {
			os.Exit(0)
		}
		i++
	}
}
