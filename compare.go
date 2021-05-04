package piscine

func Compare(a, b string) int {
	str1 := []rune(a)
	str2 := []rune(b)
	check1 := 0
	if len(a) >= len(b) {
		check1 = len(b)
	} else {
		check1 = len(a)
	}
	for i := 0; i < check1; i++ {
		if str1[i] > str2[i] {
			return 1
		} else if str2[i] > str1[i] {
			return -1
		}
	}
	return 0
}
