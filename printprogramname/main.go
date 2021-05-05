package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	prog_name := os.Args[0]
	prog_name = string(prog_name)
	for _, word := range prog_name[2:] {
		z01.PrintRune(word)
	}
}
