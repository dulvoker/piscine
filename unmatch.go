package piscine

func Unmatch(a []int) int {
	check := 0
	for i := 0; i < len(a); i++ {
		check = 0
		for q := 0; q < len(a); q++ {
			if a[i] == a[q] && q != i {
				check = 1
			}
		}
		if check == 0 {
			return a[i]
		}
	}
	return -1
}
