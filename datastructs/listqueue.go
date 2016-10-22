package datastructs

type ListQueue struct {
	list  *Slist
	count int
}

func NewListQueue() *ListQueue {
	q := &ListQueue{list: &Slist{len: 0}}
	return q
}

func (q *ListQueue) Enqueue(v interface{}) {
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
