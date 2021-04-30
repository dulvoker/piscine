package piscine

func UltimateDivMod(a *int, b *int) {
	c := *b
	y := *a
	*a = c / y
	*b = c % y
}
