package main

import "github.com/01-edu/z01"


func main() {

	for r := 97; r < 123; r++ {
		z01.PrintRune(rune(r))
		if r < 122 {
			z01.PrintRune('\n')
		}
	}
}