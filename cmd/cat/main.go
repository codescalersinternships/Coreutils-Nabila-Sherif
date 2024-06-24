package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic((e))
	}
}
func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		data, err := os.ReadFile(arg)
		check(err)
		fmt.Println(string(data))
	} else { //log.fatal or os.exit(1)
		os.Exit(1)
	}
}
