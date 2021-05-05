package piscine

func BasicJoin(elems []string) string {
	chetam := ""
	for _, word := range elems {
		chetam = chetam + word
	}
	return chetam
}
