package piscine

func IsPrintable(s string) bool {
	casted := []rune(s)
	for i := 0; i < len(s); i++ {
		if casted[i] < 33 || casted[i] > 126 {
			return false
		}
	}
	return true
}
