package piscine

func FindNextPrime(nb int) int {
	check := 0
	if nb <= 2 {
		return 2
	} else if nb < 10000 {
		for z := nb; ; z++ {
			for i := 2; i < z; i++ {
				if z%i == 0 {
					check = 1
				}
			}
			if check == 0 {
				return z
			} else {
				check = 0
			}
		}
	}
	check1 := 0
	for i := 100; i <= nb; i++ {
		if i*i > nb {
			check1 = i
			break
		}
	}
	check3 := 0
	for i := nb; i <= nb+check1*check1; i++ {
		for q := 1; q <= check1; q++ {
			if i%FindNextPrime(q) != 0 {
				check3 = 1
				q = FindNextPrime(q)
			} else {
				check3 = 0
				break
			}
		}
		if check3 == 1 {
			return i
		}
	}
	return 0
}
