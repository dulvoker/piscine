package piscine

func ListSize(l *List) int {
	q := 0
	var lis *NodeL
	lis = l.Head
	for lis != nil {
		q++
		lis = lis.Next
	}
	return q
}
