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
	lis.Data = table[0]
	var node *NodeI
	for _, each := range table[1:] {
		node = &NodeI{Data: each}
		l.Next = node
		l = l.Next
	}
	return l
}
