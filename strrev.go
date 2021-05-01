package piscine

func StrRev(s string) string {
	reversenoyet := []byte(s)
	for i := 0; i < len(s); i++ {
		reversenoyet[i] = s[len(s)-i -1]
	}
	return string(reversenoyet)
}
