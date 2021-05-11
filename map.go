package piscine

func Map(f func(int) bool, a []int) []bool {
	var bools []bool
	for _, each := range a {
		bools = append(bools, f(each))
	}
	return bools
}
