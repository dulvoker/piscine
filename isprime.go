package piscine

func IsPrime(nb int) bool {
	if nb == 1 || nb == 0 {
		return false
	} else if nb == 2 {
		return true
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
