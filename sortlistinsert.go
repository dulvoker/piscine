package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	ListSort(l)
	l = listPushBack(l, data_ref)
	ListSort(l)
	return l
}
