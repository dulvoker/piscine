package piscine

func SortTable(table []int) {
	len := len(table)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

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
	SortTable(table)
	lis = nil
	for i := 0; i < len(table); i++ {
		lis = listPushBack(lis, table[len(table)-1-i])
	}
	return l
}
