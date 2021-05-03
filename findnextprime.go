package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	check := 0
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			check = 1
		}
	}
	if check == 0 {
		return nb
	} else {
		check = 0
		for z := nb + 1; ; z++ {
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
	return 0
}
