package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, word := range os.Args[1:] {
		for _, inword := range word {
			z01.PrintRune(inword)
		}
		z01.PrintRune('\n')
	}
}
