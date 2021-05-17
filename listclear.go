package piscine

func ListClear(l *List) {
	for l.Head != l.Tail {
		l.Head = nil
		l.Head = l.Head.Next
	}
}
