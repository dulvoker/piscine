package piscine

func IsNumeric(s string) bool {
	casted := []rune(s)
	for i := 0; i < len(s); i++ {
		if casted[i] < 48 || casted[i] > 57 {
			return false
		}
	}
	return true
}
