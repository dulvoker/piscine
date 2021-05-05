package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	word := ""
	for i := 1; i < len(os.Args); i++ {
		word = os.Args[len(os.Args)-i]
		for _, inword := range word {
			z01.PrintRune(inword)
		}
		z01.PrintRune('\n')
	}
}
