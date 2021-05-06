package main

import (
	"os"

	"github.com/01-edu/z01"
)

func calc(s string) rune {
	sam := []rune(s)
	prev := 0
	for i := 0; i < len(s); i++ {
		if sam[i] > 47 && sam[i] < 58 {
			prev = prev * 10
			prev = int(sam[i]) - 48 + prev
		}
	}
	if prev <= 0 || prev > 25 {
		return rune(32)
	}
	return rune(prev + 48 + 48)
}

func upper(s string) rune {
	sam := []rune(s)
	prev := 0
	for i := 0; i < len(s); i++ {
		if sam[i] > 47 && sam[i] < 58 {
			prev = prev * 10
			prev = int(sam[i]) - 48 + prev
		}
	}
	if prev <= 0 || prev > 25 {
		return rune(32)
	}
	return rune(prev + 48 + 16)
}

func main() {
	if len(os.Args) == 1 {
	} else if os.Args[1] == "--upper" {
		for _, word := range os.Args[2:] {
			z01.PrintRune(upper(word))
		}
		z01.PrintRune('\n')
	} else {
		for _, word := range os.Args[1:] {
			z01.PrintRune(calc(word))
		}
		z01.PrintRune('\n')
	}
}
