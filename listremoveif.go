package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	var table []*NodeL
	var lis *NodeL
	lis = l.Head
	for lis != nil {
		if lis.Data != data_ref {
			table = append(table, lis)
		}
		lis = lis.Next
	}
	ListClear(l)
	for _, each := range table {
		ListPushBack(l, each.Data)
	}
}
