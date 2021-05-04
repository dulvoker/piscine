package piscine

func TrimAtoi(s string) int {
	casted := []rune(s)
	firstnumbcheck := 0
	var table []rune
	signcheck := 0
	for i := 0; i < len(s); i++ {
		if casted[i] > 47 && casted[i] < 58 {
			if firstnumbcheck == 1 {
				table = append(table, rune(casted[i]))
			} else if casted[i] != 47 {
				table = append(table, rune(casted[i]))
				firstnumbcheck = 1
			}
		}
		if casted[i] == 45 && len(table) == 0 {
			signcheck = 1
		}
	}
	sum := 0
	tens := 1
	for i := 0; i < len(table); i++ {
		for z := 0; z < len(table)-i-1; z++ {
			tens = tens * 10
		}
		sum = sum + (int(table[i])-48)*tens
		tens = 1
	}
	if signcheck == 1 {
		return -sum
	}
	return sum
}
