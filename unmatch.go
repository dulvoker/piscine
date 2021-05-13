package piscine

func Unmatch(a []int) int {
	firstnumb := a[0]
	var pattern []int
	for i, numb := range a {
		if i != 0 && numb == firstnumb {
			break
		}
		pattern = append(pattern, numb)
	}
	for i, numb := range a {
		if pattern[i%len(pattern)] != numb {
			return numb
		}
	}
	return -1
}
