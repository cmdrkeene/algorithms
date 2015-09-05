package algorithms

func NewQueue() *Queue {
	return &Queue{list: NewLinkedList()}
}

type Queue struct {
	list *LinkedListNode
}

func (self *Queue) Enqueue(elements ...interface{}) {
	for _, e := range elements {
		self.list = self.list.Append(e)
	}
}

func (self *Queue) Dequeue() interface{} {
	if self.list.Empty() {
		return nil
	}

	head := self.list.Data()
	self.list = self.list.Remove(head)
	return head
}

func (self *Queue) Empty() bool {
	return self.list.Empty()
}

func (self *Queue) Size() int {
	return self.list.Size()
}

func (self *Queue) Peek() interface{} {
	return self.list.Data()
}
