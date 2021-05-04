package piscine

func NRune(s string, n int) rune {
	abc := []rune(s)
	if n > len(s) || n <= 0 {
		return 0
	}
	return abc[n-1]
}
