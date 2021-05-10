package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintNbr(n int) string {
	var table []rune
	for i := 0; ; i++ {
		table = append(table, rune(n%10+48))
		if n/10 == 0 {
			break
		}
		n = n / 10
	}
	newstring := ""
	lenq := len(table)
	for i := 0; i < lenq; i++ {
		newstring = newstring + string(table[lenq-i-1])
	}
	return newstring
}

func main() {
	points := point{}

	setPoint(&points)

	for_print := "x = " + PrintNbr(points.x) + ", y = " + PrintNbr(points.y) + "\n"

	for _, char := range for_print {
		z01.PrintRune(char)
	}
}
