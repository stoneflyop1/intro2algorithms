package datastructs

type ListStack struct {
	list  *Slist
	len   int
	count int
}

func NewListStack(capacity int) *ListStack {
	return &ListStack{len: capacity, list: &Slist{len: 0}}
}

func (s *ListStack) IsEmpty() bool {
	return s.list.root == nil
}

func (s *ListStack) Push(v interface{}) {
	if s.count == s.len {
		panic("overflow")
	}
	s.list.Append(v)
	s.count++
}

func (s *ListStack) Pop() interface{} {
	if s.list.root == nil {
		panic("underflow")
	}
	last, _ := s.list.Last()
	s.list.remove(last)
	s.count--
	return last.value
}
