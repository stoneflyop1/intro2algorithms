package datastructs

type ListStack struct {
	list *Slist
	//len   int
	count int
}

func NewListStack() *ListStack {
	return &ListStack{list: &Slist{len: 0}}
}

func (s *ListStack) IsEmpty() bool {
	return s.list.root == nil
}

func (s *ListStack) Push(v interface{}) {
	s.list.AppendFront(v)
	s.count++
}

func (s *ListStack) Pop() interface{} {
	if s.list.root == nil {
		panic("underflow")
	}
	node := s.list.RemoveFront()
	s.count--
	return node.value
}
