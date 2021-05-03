package piscine

func fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	} else if index == 1 || index == 2 {
		return 1
	} else if index == 3 {
		return 2
	} else {
		return fibonacci(index-1) + fibonacci(index-2)
	}
}
