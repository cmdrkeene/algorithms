package algorithms

type Queue interface {
	Enqueue(...interface{})
	Dequeue() interface{}
	Peek() interface{}
	Size() int
	Empty() bool
}

func NewStackQueue() *StackQueue {
	return &StackQueue{
		oldest: NewStack(),
		newest: NewStack(),
	}
}

type StackQueue struct {
	oldest *Stack
	newest *Stack
}

func (self *StackQueue) Enqueue(elements ...interface{}) {
	self.newest.Push(elements...)
}

func (self *StackQueue) shift() {
	if self.oldest.Empty() {
		for !self.newest.Empty() {
			self.oldest.Push(self.newest.Pop())
		}
	}
}

func (self *StackQueue) Dequeue() interface{} {
	self.shift()
	return self.oldest.Pop()
}

func (self *StackQueue) Peek() interface{} {
	self.shift()
	return self.oldest.Peek()
}

func (self *StackQueue) Size() int {
	return self.newest.Size() + self.oldest.Size()
}

func (self *StackQueue) Empty() bool {
	return self.Size() == 0
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{list: NewLinkedList()}
}

type LinkedListQueue struct {
	list *LinkedListNode
}

func (self *LinkedListQueue) Enqueue(elements ...interface{}) {
	for _, e := range elements {
		self.list = self.list.Append(e)
	}
}

func (self *LinkedListQueue) Dequeue() interface{} {
	if self.list.Empty() {
		return nil
	}

	head := self.list.Data()
	self.list = self.list.Remove(head)
	return head
}

func (self *LinkedListQueue) Empty() bool {
	return self.list.Empty()
}

func (self *LinkedListQueue) Size() int {
	return self.list.Size()
}

func (self *LinkedListQueue) Peek() interface{} {
	return self.list.Data()
}
