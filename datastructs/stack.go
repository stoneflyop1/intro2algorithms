package datastructs

type Stack struct {
	array   []interface{}
	top     int
	fixSize int
}

func NewStack(capacity int) *Stack {
	s := new(Stack)
	if capacity <= 0 {
		panic("Stack must have Capacity")
	}
	s.array = make([]interface{}, capacity)
	s.fixSize = capacity
	s.top = -1
	return s
}

func (s *Stack) IsEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}

func (s *Stack) Push(x interface{}) {
	if s.top == s.fixSize {
		panic("overflow")
	}
	s.top = s.top + 1
	s.array[s.top] = x
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		panic("underflow")
	}
	x := s.array[s.top]
	s.top = s.top - 1
	return x
}
