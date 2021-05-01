package piscine

func SortIntegerTable(table []int) {
	len := len(table)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
