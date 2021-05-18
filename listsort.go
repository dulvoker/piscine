package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func NodePushBack(l *NodeI, data int) {
	node := &NodeI{Data: data}
	l.Next = node
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
	l.Data = table[0]
	for _, each := range table[1:] {
		NodePushBack(l, each)
		l = l.Next
	}
	return l
}
