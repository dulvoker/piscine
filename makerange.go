package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	table := make([]int, max-min)
	for i := min; i < max; i++ {
		table[i-min] = i
	}
	return table
}
