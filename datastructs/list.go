package datastructs

import "fmt"

// ref https://golang.org/src/container/list/list.go
type Slist struct {
	root *Snode
	end  *Snode
	len  int
}

type Snode struct {
	value interface{}
	next  *Snode
}

func (l *Slist) append(node *Snode) *Snode {
	l.len++
	if l.root == nil {
		l.root = node
	} else {
		if l.end == nil {
			l.end = node
			l.root.next = l.end
		} else {
			l.end.next = node
			l.end = node
		}
	}
	return node
}

func (l *Slist) Append(v interface{}) *Snode {
	node := &Snode{value: v}
	return l.append(node)
}

func (l *Slist) appendFront(node *Snode) *Snode {
	l.len++
	if l.root == nil {
		l.root = node
		return node
	}
	node.next = l.root
	l.root = node
	return node
}

func (l *Slist) AppendFront(v interface{}) *Snode {
	node := &Snode{value: v}
	return l.appendFront(node)
}

func (l *Slist) RemoveFront() *Snode {
	root := l.root
	l.root = l.root.next
	l.len--
	return root
}

func (l *Slist) print() {
	s := make([]interface{}, 0)
	s = append(s, l.len)
	node := l.root
	for node != nil {
		s = append(s, node.value)
		node = node.next
	}
	fmt.Println(s...)
}
