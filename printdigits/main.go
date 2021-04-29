package main

import "github.com/01-edu/z01"

func main() {
	for r := 48; r < 58; r++ {
		z01.PrintRune(rune(r))
	}
	z01.PrintRune('\n')
}
