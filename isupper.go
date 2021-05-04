package piscine

func IsUpper(s string) bool {
	casted := []rune(s)
	for i := 0; i < len(s); i++ {
		if casted[i] < 65 || casted[i] > 90 {
			return false
		}
	}
	return true
}
