package piscine

func UltimateDivMod(a *int, b *int) {
	temp1 := *a
	temp2 := *b
	*a = temp1 / temp2
	*b = temp1 % temp2
}
