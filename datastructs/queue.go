package datastructs

type Queue struct {
	array      []interface{}
	tail, head int
	fixSize    int
}

func NewQueue(capacity int) *Queue {
	q := new(Queue)
	if capacity <= 0 {
		panic("Queue must have Capacity")
	}
	q.array = make([]interface{}, capacity)
	q.fixSize = capacity
	return q
}

func (q *Queue) Enqueue(x interface{}) {
	if q.tail == q.fixSize {
		panic("overflow")
	}
	q.array[q.tail] = x
	q.tail = q.tail + 1
}

func (q *Queue) Dequeue() interface{} {
	if q.head == q.fixSize {
		panic("underflow")
	}
	x := q.array[q.head]
	q.head = q.head + 1
	return x
}
