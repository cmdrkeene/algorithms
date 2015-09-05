package algorithms

func NewStack() *Stack {
	return &Stack{
		list: NewLinkedList(),
	}
}

type Stack struct {
	list *LinkedListNode
}

func (s *Stack) Peek() interface{} {
	return s.list.Tail().Data()
}

func (s *Stack) Push(elements ...interface{}) {
	for _, e := range elements {
		s.list = s.list.Append(e)
	}
}

func (s *Stack) Pop() interface{} {
	if s.list.Empty() {
		return nil
	}

	tail := s.list.Tail().Data()
	s.list = s.list.Remove(tail)
	return tail
}

func (s *Stack) Size() int {
	return s.list.Size()
}

func (s *Stack) String() string {
	return s.list.String()
}
