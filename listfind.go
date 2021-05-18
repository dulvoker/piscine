package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	var lis *NodeL
	lis = l.Head
	for lis != nil {
		if comp(lis.Data, ref) {
			return &lis.Data
		}
		lis = lis.Next
	}
	return nil
}
