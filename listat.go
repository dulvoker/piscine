package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	q := 0
	for q != pos {
		if l == nil {
			return nil
		}
		l = l.Next
		q++
	}
	return l
}
