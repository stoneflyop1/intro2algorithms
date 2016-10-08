package datastructs

// Deque array implemented linked list
type Deque struct {
	array       []interface{}
	left, right int
	fixSize     int
}

func NewDeque(capacity int) *Deque {
	d := new(Deque)
	if capacity <= 0 {
		panic("deque must have capacity")
	}
	d.array = make([]interface{}, capacity)
	d.fixSize = capacity
	return d
}

func (q *Deque) EnqueueLeft(x interface{}) {

	if q.left == q.right && q.left != 0 {
		panic("overflow")
	}
	q.array[q.left] = x
	q.left = q.left - 1
	if q.left == -1 {
		q.left = q.fixSize - 1
	}
}

func (q *Deque) EnqueueRight(x interface{}) {
	if q.right == q.fixSize {
		panic("overflow")
	}
	q.array[q.right] = x
	q.right = q.right + 1
}

func (q *Deque) DequeueLeft() interface{} {
	if q.left == q.fixSize {
		panic("overflow")
	}
	q.left = q.left + 1
	if q.left == q.fixSize {
		q.left = 0
	}
	x := q.array[q.left]
	return x
}

func (q *Deque) DequeueRight() interface{} {
	if q.right == q.left {
		panic("underflow")
	}
	q.right = q.right - 1
	if q.right == -1 {
		q.right = q.fixSize - 1
	}
	x := q.array[q.right]
	return x
}
