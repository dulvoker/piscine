package piscine

func BasicAtoi2(s string) int {
	sam := []rune(s)
	prev := 0
	for i := 0; i < len(s); i++ {
		if sam[i] > 47 && sam[i] < 58 {
			prev = prev * 10
			prev = int(sam[i]) - 48 + prev
		} else return 0
	}
	return prev
}
