package piscine

func atoi(s string) int {
	sam := []rune(s)
	prev := 0
	check := 0
	checkminus := 0
	for i := 0; i < len(s); i++ {
		if sam[i] > 47 && sam[i] < 58 || sam[i] == 43 || sam[i] == 45 {
			if sam[i] == 45 || sam[i] == 43 {
				if sam[i] == 43 {
					checkminus = 1
				}
				check = 1 + check
				if check > 1 {
					return 0
				}
			}
			prev = prev * 10
			prev = int(sam[i]) - 48 + prev
		} else {
			return 0
		}
	}
	if checkminus == 1 {
		return -prev
	}
	return prev
}
