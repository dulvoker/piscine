package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	prog_name := string(os.Args[0])
	prog := []rune(prog_name)
	for _, word := range prog[2:] {
		z01.PrintRune(word)
	}
}
