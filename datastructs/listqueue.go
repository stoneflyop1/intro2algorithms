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
	q.list.Append(v)
}

func (q *ListQueue) Dequeue() interface{} {
	if q.count == 0 {
		panic("underflow")
	}
	node := q.list.RemoveFront()
	q.count--
	return node.value
}
