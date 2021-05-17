package piscine

func ListSize(l *List) int {
	q := 0
	for l.Head != l.Tail {
		q++
		l.Head = l.Head.Next
	}
	return q
}
