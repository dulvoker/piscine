package piscine

func SplitWhiteSpaces(s string) []string {
	word := ""
	var words []string
	for i, char := range s {
		if i == len(s)-1 {
			words = append(words, word)
		}
		if char != rune(32) && char != rune(9) && char != rune(10) {
			word = word + string(char)
		} else {
			words = append(words, word)
			word = ""
		}
	}
	return words
}
