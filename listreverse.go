package piscine

func ListReverse(l *List) {
	var table []*NodeL
	var lis *NodeL
	lis = l.Head
	for lis != nil {
		table = append(table, lis)
		lis = lis.Next
	}
	ListClear(l)
	for _, each := range table {
		ListPushFront(l, each.Data)
	}
}
