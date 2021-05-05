package piscine

func Join(strs []string, sep string) string {
	chetam := ""
	for _, word := range strs {
		chetam = chetam + word + sep
	}
	return chetam
}
