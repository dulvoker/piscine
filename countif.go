package piscine

func CountIf(f func(string) bool, tab []string) int {
	total := 0
	for _, each := range tab {
		if f(each) {
			total++
		}
	}
	return total
}
