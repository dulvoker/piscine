package piscine

func Compact(ptr *[]string) int {
	sum := 0
	var table []string
	for _, char := range *ptr {
		if len(char) != 0 {
			sum++
			table = append(table, char)
		}
	}
	for i := 0; i < len(*ptr); i++ {
		if i < sum {
			(*ptr)[i] = table[i]
		} else {
			(*ptr)[i] = string(rune(0))
		}
	}

	(*ptr) = (*ptr)[:sum]
	return sum
}
