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
		s.list = s.list.Add(e)
	}
}

func (s *Stack) Pop() interface{} {
	if s.list.Size() == 0 {
		return nil
	}

	data := s.list.Tail().Data()
	s.list = s.list.Remove(data)
	return data
}

func (s *Stack) Size() int {
	return s.list.Size()
}
