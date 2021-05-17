package piscine

func ListLast(l *List) interface{} {
	if ListSize(l) > 0 {
		if ListSize(l) == 1 {
			return l.Head.Data
		}
		return l.Tail.Data
	} else {
		return nil
	}
}
