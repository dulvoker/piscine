package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, word := range a {
		for _, chars := range word {
			z01.PrintRune(chars)
		}
		z01.PrintRune('\n')
	}
}
