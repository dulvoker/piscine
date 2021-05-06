package piscine

func ConcatParams(args []string) string {
	str := ""
	for _, word := range args {
		str = str + word + "\n"
	}
	return str[:len(str)-1]
}
