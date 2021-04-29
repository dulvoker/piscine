package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for r := 48; r < 56; r++ {
		for i := r + 1; i < 58; i++ {
			for z := i + 1; z < 58; z++ {
				z01.PrintRune(rune(r))
				z01.PrintRune(rune(i))
				z01.PrintRune(rune(z))
				if r == 55 && i == 56 && z == 57 {
					z01.PrintRune('\n')
				} else {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
