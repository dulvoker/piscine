package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	var table []int
	var lis *NodeI
	lis = l
	for lis != nil {
		table = append(table, lis.Data)
		lis = lis.Next
	}
	SortIntegerTable(table)
	l = nil
	for _, each := range table {
		l.Data = each
		l = l.Next
	}
	return l
}
