package piscine

func Index(s string, toFind string) int {
	castedone := []rune(s)
	castedsub := []rune(toFind)
	coord := 0
	for i := 0; i < len(s); i++ {
		if castedone[i] == castedsub[coord] {
			coord = coord + 1
			if coord == len(castedsub) {
				return (i - coord + 1)
			}
		}
	}
	return -1
}
