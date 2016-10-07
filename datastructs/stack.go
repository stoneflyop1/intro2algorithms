package datastructs

type Stack struct {
	array   []interface{}
	top     int
	fixSize int
}

func NewStack(capacity int) *Stack {
	s := new(Stack)
	if capacity < 0 {
		capacity = 0
	}
	if capacity == 0 {
		panic("Stack must have Capacity")
	}
	s.array = make([]interface{}, capacity)
	s.fixSize = capacity

	return s
}

func (s *Stack) IsEmpty() bool {
	if s.top == 0 {
		return true
	}
	return false
}

func (s *Stack) Push(x interface{}) {
	if s.fixSize != 0 {
		if s.top == s.fixSize-1 {
			panic("overflow")
		}
		s.top = s.top + 1
		s.array[s.top] = x
	} else {
		s.array = append(s.array, x)
		s.top = s.top + 1
	}
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		panic("underflow")
	}
	s.top = s.top - 1
	return s.array[s.top+1]
}
