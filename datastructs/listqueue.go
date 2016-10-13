package datastructs

type ListQueue struct {
	list  *Slist
	len   int
	count int
}

func NewListQueue(capacity int) *ListQueue {
	q := &ListQueue{len: capacity, list: &Slist{len: 0}}
	return q
}

func (q *ListQueue) Enqueue(v interface{}) {
	if q.count == q.len {
		panic("overflow")
	}
	q.count++
	q.list.AppendFront(v)
}

func (q *ListQueue) Dequeue() interface{} {
	if q.count == 0 {
		panic("underflow")
	}
	last, _ := q.list.Last()
	node := q.list.remove(last)
	q.count--
	return node.value
}
