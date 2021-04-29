package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for r := 48; r < 58; r++ {
		z01.PrintRune(rune(r))
		for i := r + 1; i < 58; i++ {
			z01.PrintRune(rune(i))
			for z := r + 1; z < 58; z++ {
				z01.PrintRune(rune(z))
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
