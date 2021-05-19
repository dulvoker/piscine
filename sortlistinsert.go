package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	var table []int
	var lis *NodeI
	lis = l
	for lis != nil {
		table = append(table, lis.Data)
		lis = lis.Next
	}
	table = append(table, data_ref)
	SortTable(table)
	lis = nil
	for _, each := range table {
		lis = listPushBack(lis, each)
	}
	return lis
}
