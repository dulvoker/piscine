package piscine

func Index2(s string, toFind string) []int {
	castedone := []rune(s)
	castedsub := []rune(toFind)
	coord := 0
	z := len(s)
	q := len(toFind)
	if q == 0 {
		return nil
	}
	var coords []int
	for i := 0; i < z; i++ {
		if castedone[i] == castedsub[coord] {
			coord = coord + 1
			if coord == q {
				coords = append(coords, i-1)
				coord = 0
			}
		}
	}
	if len(coords) == 0 {
		return nil
	}
	return coords
}

func Split(s, sep string) []string {
	numbers := Index2(s, sep)
	var newstring []string
	word := ""
	for _, char := range s {
		word = word + string(char)
	}
	temp := 0
	q := 0
	for i := 0; i < len(numbers); i++ {
		temp = numbers[i]
		newstring = append(newstring, word[q:temp])
		q = temp + len(sep)
	}
	newstring = append(newstring, word[q:])
	return newstring
}
