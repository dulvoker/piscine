package piscine

func UltimateDivMod(a *int, b *int) {
	c := *a
	y := *b
	*a = c / y
	*b = c % y
}
