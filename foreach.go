package piscine

import "github.com/01-edu/z01"

func PrintNbr(b int) {
	z01.PrintRune(rune(b + 48))
}

func ForEach(f func(int), numbers []int) {
	for _, each := range numbers {
		PrintNbr(each)
	}
}
