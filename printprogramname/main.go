package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	prog_name := os.Args[0]
	prog_name = string(prog_name)
	casted := []rune(prog_name)
	for i := 2; i < len(prog_name); i++ {
		z01.PrintRune(casted[i])
	}
}
