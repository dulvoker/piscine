package piscine

func IsAlpha(s string) bool {
	casted := []rune(s)
	for i := 0; i < len(s); i++ {
		if casted[i] < 48 || casted[i] > 57 && casted[i] < 65 || casted[i] > 90 && casted[i] < 97 || casted[i] > 122 {
			return false
		}
	}
	return true
}
