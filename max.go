package piscine

func Max(a []int) int {
	maxnumb := 0
	for _, each := range a {
		if each >= maxnumb {
			maxnumb = each
		}
	}
	return maxnumb
}
