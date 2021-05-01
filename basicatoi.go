package piscine

func BasicAtoi(s string) int {
	sam := []rune(s)
	summa := 0
	z := 0
	for i := 0; i < len(s); i++ {
		if sam[i] > 48 && sam[i] < 58 {
			for q := 0; q < z; q++ {
				summa = summa * 10
			}
			z = z + 1
			summa = ((sam[i] - '0') - 48 + summa)
		}
	}
	return sum
}
