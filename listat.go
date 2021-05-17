package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	q := 0
	for q != pos {
		l = l.Next
		if l == nil {
			return nil
		}
		q++
	}
	return l
}
