package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func printDirectory(directoryPath string, depth int, currentDepth int) {
	if depth == currentDepth {
		return
	}
	d, err := os.ReadDir(directoryPath)
	check(err)
	for _, currentLoc := range d {
		for i := 0; i < currentDepth; i++ {
			fmt.Print("   ")
		}
		fmt.Println("|―", currentLoc.Name())
		if currentLoc.IsDir() {
			printDirectory(filepath.Join(directoryPath, currentLoc.Name()), depth, currentDepth+1)
		}
	}
}
func main() {
	var depth int
	flag.IntVar(&depth, "L", 1, "for indicating depth")
	flag.Parse()
	c := flag.Args()[0]
	printDirectory(c, depth, 0) //recursion till end of depth
}
