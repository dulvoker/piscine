package piscine

func Index(s string, toFind string) int {
	castedone := []rune(s)
	castedsub := []rune(toFind)
	coord := 0
	z := len(s)
	q := len(toFind)
	if q == 0 {
		return 0
	}
	for i := 0; i < z; i++ {
		if castedone[i] == castedsub[coord] {
			coord = coord + 1
			if coord == q {
				return (i - coord + 1)
			}
		}
	}
	return -1
}
