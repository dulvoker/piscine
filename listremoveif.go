package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	var lis *NodeL
	lis = l.Head
	for lis != nil {
		if CompStr(lis.Next.Data, data_ref) {
			lis.Next = lis.Next.Next
		}
		lis = lis.Next
	}
}
