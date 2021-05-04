package piscine

func ToUpper(s string) string {
	casted := []rune(s)
	for i := 0; i < len(s); i++ {
		if casted[i] > 96 && casted[i] < 123 {
			casted[i] = rune(casted[i] - 32)
		}
	}
	return string(casted)
}
