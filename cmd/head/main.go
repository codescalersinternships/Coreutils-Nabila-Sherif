package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic((e))
	}
}
func main() {
	nPtr := flag.Int("n", 10, "an int")
	flag.Parse()
	arg := os.Args[1]
	if strings.Contains(arg, "-") {
		arg = os.Args[2]
	}
	data, e := os.ReadFile(arg)
	check(e)
	var s []string
	s = strings.Split(string(data), "\n")
	for j := 0; j < len(s) && j < *nPtr; j++ {
		if s[j] != "" {
			fmt.Println(s[j])
		} else {
			os.Exit(0)
		}
	}
}
