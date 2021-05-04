package piscine

func IsLower(s string) bool {
	casted := []rune(s)
	for i := 0; i < len(s); i++ {
		if casted[i] < 97 || casted[i] > 122 {
			return false
		}
	}
	return true
}
