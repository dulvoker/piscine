package piscine

func Concat(str1 string, str2 string) string {
	str3 := []rune(str1 + str2)
	return string(str3)
}
