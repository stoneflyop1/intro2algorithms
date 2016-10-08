package datastructs

import "fmt"

func Tests() {
	fmt.Println("stack test...")

	s := NewStack(4)
	fmt.Println(s.IsEmpty())
	s.Push(1)
	fmt.Println(s.IsEmpty())
	x := s.Pop()
	fmt.Println(x == 1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	//s.Push(5)

	q := NewQueue(4)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	//q.Enqueue(5)
	y := q.Dequeue()
	fmt.Println(y == 1)
	q.Dequeue()
	q.Dequeue()

	q0 := NewDeque(4)
	q0.EnqueueLeft(1)
	q0.EnqueueLeft(2)
	q0.EnqueueRight(3)
	q0.EnqueueRight(4)
	y0 := q0.DequeueLeft()
	fmt.Println(y0 == 2)
	q0.DequeueRight()
	q0.DequeueLeft()
	q0.DequeueRight()
	q0.DequeueRight()
}
