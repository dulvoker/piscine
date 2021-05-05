package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	prog_name := string(os.Args[0])
	for _, word := range prog_name[2:] {
		z01.PrintRune(word)
	}
}
