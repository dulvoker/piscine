package main

func IterativeFactorial(nb int) int {
	if nb > 1 {
		return nb * IterativeFactorial(nb-1)
	} else if nb == 1 || nb == 0 {
		return 1
	} else {
		return 0
	}

}
