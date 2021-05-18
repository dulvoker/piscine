package piscine

func ListForEach(l *List, f func(*NodeL)) {
	var lis *NodeL
	lis = l.Head
	for lis != nil {
		f(lis)
		lis = lis.Next
	}
}
