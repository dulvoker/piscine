package piscine

func ListLast(l *List) interface{} {
	if ListSize(l) > 0 {
		return l.Tail.Data
	} else {
		return nil
	}
}
