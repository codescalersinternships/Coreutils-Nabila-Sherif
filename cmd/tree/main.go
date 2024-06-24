package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic((e))
	}
}
func createEmptyFile(name string) {
	d := []byte("")
	check(os.WriteFile(name, d, 0644))
}
func main() {
	depth := os.Args[len(os.Args)-1]
	i, e := strconv.Atoi(depth)
	check(e)

	err := os.Mkdir("subdir", 0755)
	check(err)
	defer os.RemoveAll("subdir")
	os.MkdirAll("subdir/parent/child1", 0755)
	os.MkdirAll("subdir/parent/child2", 0755)
	os.MkdirAll("subdir/parent/child3", 0755)
	//check(err)
	createEmptyFile("subdir/parent/child1/file2")
	createEmptyFile("subdir/parent/child2/file3")
	createEmptyFile("subdir/parent/child3/file4")
	c, err := os.ReadDir("subdir/parent")
	check(err)
	for _, entry := range c {
		fmt.Println("|")
		fmt.Println("――", entry.Name())
		if i == 2 {
			if entry.IsDir() {
				d, err := os.ReadDir("subdir/parent/" + entry.Name())
				check(err)
				for _, ent := range d {
					fmt.Println("    |")
					fmt.Println("    ――", ent.Name())
				}
			}
		}

	}
}
