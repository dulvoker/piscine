package main

func BasicAtoi(s string) int {
	sam := []rune(s)
	prev := 0
	for i := 0; i < len(s); i++ {
		if sam[i] > 48 && sam[i] < 58 {
			prev = prev * 10
			prev = int(sam[i]) - 48 + prev
		}
	}
	return prev
}
