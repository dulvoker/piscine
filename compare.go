package piscine

func Compare(a, b string) int {
	if len(a) > len(b) {
		return 1
	} else if len(b) > len(a) {
		return -1
	} else {
		sum1 := 0
		sum2 := 0
		check1 := 0
		str1 := []rune(a)
		str2 := []rune(b)
		for i := 0; i < len(a); i++ {
			if str1[i] != str2[i] {
				check1 = 1
			}
		}
		if check1 == 0 {
			return 0
		} else if sum1 > sum2 {
			return 1
		} else {
			return -1
		}
	}
}
