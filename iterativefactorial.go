package main

func IterativeFactorial(nb int) int {
	sum := 1
	if nb >= 1 {
		for i := 1; i <= nb; i++ {
			sum = sum * i
		}
		return sum
	} else {
		if nb == 0 {
			return 1
		}
		return 0
	}
}
