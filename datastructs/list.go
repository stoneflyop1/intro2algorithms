package datastructs

// ref https://golang.org/src/container/list/list.go
type Slist struct {
	root *Snode
	len  int
}

type Snode struct {
	value interface{}
	next  *Snode
}

func (l *Slist) append(node *Snode) *Snode {
	root := l.root
	l.len++
	if root == nil {
		l.root = node
		return node
	}
	for root.next != nil {
		root = root.next
	}
	root.next = node
	return node
}

func (l *Slist) Append(v interface{}) *Snode {
	node := &Snode{value: v}
	return l.append(node)
}

func (l *Slist) Last() (*Snode, int) {
	root := l.root
	count := 0
	for root != nil && root.next != nil {
		root = root.next
		count++
	}
	return root, count
}

func (l *Slist) remove(node *Snode) *Snode {
	prev := l.root
	if prev == nil {
		return nil
	}
	if prev.next == nil && prev == node {
		l.root = nil
		l.len--
		return node
	}
	for prev != nil {
		if prev.next == node {
			prev.next = prev.next.next
			l.len--
			return node
		}
		prev = prev.next
	}
	return nil
}

func (l *Slist) Remove(v interface{}) *Snode {
	prev := l.root
	if prev == nil {
		return nil
	}
	for prev.next != nil {
		root := prev.next
		if root.value == v {
			prev = root.next
			l.len--
			return root
		}
		prev = root
	}
	return nil
}
