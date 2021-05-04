package piscine

func AlphaCount(s string) int {
	casted := []rune(s)
	sum1 := 0
	for i := 0; i < len(s); i++ {
		if casted[i] > 64 && casted[i] < 91 || casted[i] > 96 && casted[i] < 123 {
			sum1 = sum1 + 1
		}
	}
	return sum1
}
