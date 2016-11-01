package datastructs

import "fmt"

type IntComparable int

//implement Comparable interface for int
func (i IntComparable) compare(v interface{}) int {
	// here you cannot only use v.(type) because v may be typeof  IntComparable
	var val int
	switch v.(type) {
	case int:
		val = v.(int)
	case IntComparable:
		val = int(v.(IntComparable))
	}
	i1 := int(i)
	vi := val
	if i1 > vi {
		return 1
	} else if i1 < vi {
		return -1
	}
	return 0
}

func Tests() {
	fmt.Println("*************datastructs****************")
	fmt.Println("------stack test--------------")

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
	fmt.Println("------queue test--------------")
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

	fmt.Println("------deque test--------------")
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
	//q0.DequeueRight() //underflow

	fmt.Println("------liststack test--------------")
	ls := NewListStack()
	ls.Push(1)
	ls.Push(2)
	ls.Push(3)
	ls.Push(4)
	fmt.Println("Pop:", ls.Pop())
	fmt.Println(ls.list.len, ls.count)
	fmt.Println("Pop:", ls.Pop())
	fmt.Println(ls.list.len, ls.count)
	fmt.Println("Pop:", ls.Pop())
	fmt.Println(ls.list.len, ls.count)
	fmt.Println("Pop:", ls.Pop())
	fmt.Println(ls.list.len, ls.count)
	//fmt.Println("Pop:", ls.Pop()) //underflow
	ls.Push(1)
	ls.Push(2)
	ls.Push(3)
	ls.Push(4)
	fmt.Println(ls.list.len, ls.count)
	//ls.Push(5) //overflow

	fmt.Println("------listqueue test--------------")
	lq := NewListQueue()
	lq.Enqueue(1)
	lq.list.print()
	lq.Enqueue(2)
	lq.list.print()
	lq.Enqueue(3)
	lq.list.print()
	lq.Enqueue(4)
	lq.list.print()
	fmt.Println("Dequeue: ", lq.list.end.value, lq.Dequeue())
	fmt.Println("Dequeue: ", lq.list.end.value, lq.Dequeue())
	fmt.Println("Dequeue: ", lq.list.end.value, lq.Dequeue())
	fmt.Println("Dequeue: ", lq.list.end.value, lq.Dequeue())
	//lq.Dequeue() //underflow
	lq.Enqueue(1)
	lq.Enqueue(2)
	lq.Enqueue(3)
	lq.Enqueue(4)
	lq.Enqueue(5) //overflow

	fmt.Println("---------------binary tree test-----------")
	t4 := &BTreeNode{key: 4, parent: nil}
	// t2 := &BTreeNode{key: 2, parent: t4}
	// t1 := &BTreeNode{key: 1, parent: t2}
	// t3 := &BTreeNode{key: 3, parent: t2}
	// t5 := &BTreeNode{key: 5, parent: t4}
	// t4.right = t5
	// t4.left = t2
	// t2.left = t1
	// t2.right = t3
	t4 = TreeInsert(t4, IntComparable(2))
	t4 = TreeInsert(t4, IntComparable(1))
	t4 = TreeInsert(t4, IntComparable(3))
	t4 = TreeInsert(t4, IntComparable(5))

	PrintBinaryTree2(t4)
	fmt.Println()
	PrintBinaryTree(t4)
	st := TreeSearch(t4, IntComparable(3))
	fmt.Println()
	fmt.Println("Search 3 recursively: ", st.key)
	st2 := TreeSearchIterative(t4, IntComparable(3))
	fmt.Println("Search 3 iteratively: ", st2.key)
	fmt.Println("Find minimum: ", TreeMinimum(t4).key)
	fmt.Println("Find maximum: ", TreeMaximum(t4).key)
	fmt.Println("Find Successor for 3: ", TreeSuccessor(st).key)
	fmt.Println("Find Predecessor for 3: ", TreePredecessor(st).key)
	tt := TreeInsert(t4, IntComparable(7))
	PrintBinaryTree2(tt)
	fmt.Println()
	tt = TreeInsert(tt, IntComparable(6))
	PrintBinaryTree2(tt)
	fmt.Println()
	tt = TreeDelete(tt, st)
	fmt.Println("delete", st.key)
	PrintBinaryTree2(tt)
	fmt.Println()
	tt = TreeDelete(tt, TreeMinimum(tt))
	fmt.Println("delete minimum...")
	PrintBinaryTree2(tt)
	fmt.Println()
	tt = TreeDelete(tt, TreeMaximum(tt))
	fmt.Println("delete maximum...")
	PrintBinaryTree2(tt)
	fmt.Println()

	fmt.Println("---------------arbitrary tree test-----------")
	abt := &TreeNode{key: 1, parent: nil}
	at2 := &TreeNode{key: 2, parent: abt}
	at4 := &TreeNode{key: 4, parent: at2}
	at5 := &TreeNode{key: 5, parent: at2}
	at3 := &TreeNode{key: 3, parent: abt}
	abt.leftChild = at2
	abt.rightSibling = at3
	at2.leftChild = at4
	at2.rightSibling = at5
	PrintTree(abt)

}
