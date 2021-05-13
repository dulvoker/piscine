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
	(*ptr) = table
	return sum
}
