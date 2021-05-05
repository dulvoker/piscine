package piscine

func Join(strs []string, sep string) string {
	chetam := ""
	for index, word := range strs {
		if index == len(strs)-1 {
			chetam = chetam + word
		}
		chetam = chetam + word + sep
	}
	return chetam
}
