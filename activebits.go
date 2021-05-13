package piscine

func decToBinary(n int) []int {
	var binaryNum []int
	i := 0
	for {
		if n == 0 {
			break
		}
		binaryNum = append(binaryNum, n%2)
		n = n / 2
		i++
	}
	return binaryNum
}

func ActiveBits(n int) int {
	numb := decToBinary(n)
	sum := 0
	for _, num := range numb {
		if num == 1 {
			sum++
		}
	}
	return sum
}
