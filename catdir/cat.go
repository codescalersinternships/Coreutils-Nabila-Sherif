package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		data, _ := os.ReadFile(arg)
		fmt.Println(string(data))
	} else { //log.fatal or os.exit(1)
		os.Exit(1)
	}
}
