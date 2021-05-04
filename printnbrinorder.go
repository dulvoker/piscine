package piscine

func PrintNbrInOrder(n int) {
	var table []rune
	for i := 0; ; i++ {
		table = append(table, rune(n%10+48))
		if n/10 == 0 {
			break
		}
		n = n / 10
	}
	len := len(table)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
	for i := 0; i < len; i++ {
		PrintRune(table[i])
	}
}
