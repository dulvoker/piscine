package piscine

func RecursiveFactorial(nb int) int {
	sum := 1
	if nb >= 1 && nb <= 20 {
		return sum * RecursiveFactorial(nb-1)
	} else {
		if nb == 0 {
			return 1
		}
		return 0
	}
}
