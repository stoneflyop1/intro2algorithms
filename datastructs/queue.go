package datastructs

type Queue struct {
	array      []interface{}
	tail, head int
	fixSize    int
}

func NewQueue(capacity int) *Queue {
	q := new(Queue)
	if capacity < 0 {
		capacity = 0
	}
	if capacity == 0 {
		panic("Queue must have Capacity")
	}
	q.array = make([]interface{}, capacity)
	q.fixSize = capacity
	return q
}

func (q *Queue) Enqueue(x interface{}) {
	if q.fixSize != 0 {
		if q.tail == q.fixSize-1 {
			panic("overflow")
		}
		q.array[q.tail] = x
		q.tail = q.tail + 1
	} else {
		q.array = append(q.array, x)
		q.tail = q.tail + 1
	}
}

func (q *Queue) Dequeue() interface{} {
	if q.head == -1 {
		panic("underflow")
	}
	x := q.array[q.head]
	q.head = q.head - 1
	return x
}
