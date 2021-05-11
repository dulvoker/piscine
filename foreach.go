package piscine

func ForEach(f func(int), numbers []int) {
	for _, each := range numbers {
		f(each)
	}
}
