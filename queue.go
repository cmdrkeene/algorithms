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
		forward: NewStack(),
		reverse: NewStack(),
	}
}

type StackQueue struct {
	forward *Stack
	reverse *Stack
}

func (self *StackQueue) Enqueue(elements ...interface{}) {
	self.forward.Push(elements...)
}

func (self *StackQueue) swap() {
	for {
		data := self.forward.Pop()
		if data == nil {
			break
		}
		self.reverse.Push(data)
	}
}

func (self *StackQueue) restore() {
	for {
		data := self.reverse.Pop()
		if data == nil {
			return
		}
		self.forward.Push(data)
	}
}

func (self *StackQueue) Dequeue() interface{} {
	self.swap()
	data := self.reverse.Pop()
	self.restore()
	return data
}

func (self *StackQueue) Peek() interface{} {
	self.swap()
	data := self.reverse.Peek()
	self.restore()
	return data
}

func (self *StackQueue) Size() int {
	return self.forward.Size()
}

func (self *StackQueue) Empty() bool {
	return self.forward.Size() == 0
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
