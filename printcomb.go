package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for r := 0; r < 10; r++ {
		z01.PrintRune(r)
		for i := r + 1; i < 10; i++ {
			z01.PrintRune(i)
			for z := r + 1; z < 10; z++ {
				z01.PrintRune(z)
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
