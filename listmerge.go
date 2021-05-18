package piscine

func ListMerge(l1 *List, l2 *List) {
	var table []*NodeL
	var lis *NodeL
	lis = l2.Head
	for lis != nil {
		table = append(table, lis)
		lis = lis.Next
	}
	for _, each := range table {
		ListPushBack(l1, each.Data)
	}
}
