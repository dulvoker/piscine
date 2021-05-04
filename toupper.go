package piscine

func ToUpper(s string) string {
	casted := []rune(s)
	for i := 0; i < len(s); i++ {
		if casted[i] > 97 && casted[i] < 122 {
			casted[i] = rune(casted[i] - 32)
		}
	}
	return string(casted)
}
