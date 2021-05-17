package piscine

func ListPushFront(l *List, data interface{}) {
	node := &NodeL{Data: data}
	if l.Head == nil {
		l.Tail = node
	} else {
		node.Next = l.Head
	}
	l.Head = node
}
