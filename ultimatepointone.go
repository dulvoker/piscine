package main

func UltimateDivMod(a *int, b *int) {
	c := a
	*a = *a / *b
	*b = c % *b
}

import (
	"fmt"
	piscine "github.com/01-edu/z01"
)

func main() {
	a := 13
	b := 2
	piscine.UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

