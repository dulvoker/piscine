package piscine

func ToLower(s string) string {
	casted := []rune(s)
	for i := 0; i < len(s); i++ {
		if casted[i] > 64 && casted[i] < 91 {
			casted[i] = rune(casted[i] + 32)
		}
	}
	return string(casted)
}
