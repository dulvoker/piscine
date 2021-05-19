package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	var table []int
	var lis, lis2 *NodeI
	lis = n1
	for lis != nil {
		table = append(table, lis.Data)
		lis = lis.Next
	}
	lis2 = n2
	for lis2 != nil {
		table = append(table, lis2.Data)
		lis2 = lis2.Next
	}
	SortTable(table)
	lis = nil
	for _, each := range table {
		lis = listPushBack(lis, each)
	}
	return lis
}
