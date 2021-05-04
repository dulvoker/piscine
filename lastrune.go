package piscine

func LastRune(s string) rune {
	abc := []rune(s)
	return abc[len(s)-1]
}
