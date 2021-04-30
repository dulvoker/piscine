package piscine

func UltimateDivMod(a *int, b *int) {
	y := *b
	x := *a
	*a = x / y
	*b = x % y
}
