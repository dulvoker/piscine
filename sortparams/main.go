package main

import (
	"os"

	"github.com/01-edu/z01"
)

func calc(s string) int {
	casted := []rune(s)
	sum := 0
	for _, char := range casted {
		sum = sum + int(char)
	}
	return sum
}

func main() {
	var table []int
	for _, word := range os.Args[1:] {
		table = append(table, calc(word))
	}
	lenq := len(table)
	for i := 0; i < lenq-1; i++ {
		for j := 0; j < lenq-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
	check := 0
	for check < lenq {
		for _, word := range os.Args[1:] {
			if table[check] == calc(word) {
				for _, inword := range word {
					z01.PrintRune(inword)
				}
				check = check + 1
				z01.PrintRune('\n')
				break
			}
		}
	}
}
