package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func main() {
	depth := os.Args[2]
	i, e := strconv.Atoi(depth)
	check(e)
	c, err := os.ReadDir(os.Args[len(os.Args)-1])
	check(err)
	for _, entry := range c {
		fmt.Println("|")
		fmt.Println("――", entry.Name())
		if i == 2 {
			if entry.IsDir() {
				d, err := os.ReadDir(entry.Name())
				check(err)
				for _, ent := range d {
					fmt.Println("    |")
					fmt.Println("     ――", ent.Name())
				}
			}
		}

	}
}
