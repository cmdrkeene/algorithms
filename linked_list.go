package algorithms

import (
	"bytes"
	"fmt"
)

func NewLinkedList(elements ...interface{}) *LinkedListNode {
	if len(elements) == 0 {
		return &LinkedListNode{}
	} else {
		head := &LinkedListNode{data: elements[0]}
		for _, e := range elements[1:] {
			head.Add(e)
		}
		return head
	}
}

type LinkedListNode struct {
	data interface{}
	next *LinkedListNode
}

func (self *LinkedListNode) Add(data interface{}) *LinkedListNode {
	tail := NewLinkedList(data)

	// empty case
	if self.data == nil {
		return tail
	}

	// n-th case
	node := self
	for {
		if node.next == nil {
			node.next = tail
			break
		}
		node = node.next
	}

	return self
}

// Remove removes first element containing data
func (self *LinkedListNode) Remove(data interface{}) *LinkedListNode {
	// List head removed
	if self.data == data {
		return self.next
	}

	// Scan list body
	node := self
	for {
		// End of list reached
		if node == nil {
			break
		}

		// Look ahead to next and skip if matched
		if node.next != nil && node.next.data == data {
			node.next = node.next.next
			break
		}

		// Advance cursor
		node = node.next
	}

	return self
}

func (self *LinkedListNode) Empty() bool {
	return self.Size() == 0
}

func (self *LinkedListNode) Size() int {
	var size int
	node := self
	for {
		if node == nil {
			break
		}
		size++
		node = node.next
	}

	return size
}

func (self *LinkedListNode) String() string {
	buf := bytes.NewBufferString("")

	node := self
	for {
		if node == nil {
			break
		}
		fmt.Fprintf(buf, "%v", node.data)

		if node.next != nil {
			buf.Write([]byte("->"))
		}

		node = node.next
	}
	return buf.String()
}

func (self *LinkedListNode) Head() *LinkedListNode {
	return self
}

func (self *LinkedListNode) Tail() *LinkedListNode {
	node := self
	for {
		if node.next == nil {
			break
		}
		node = node.next
	}
	return node
}

func (self *LinkedListNode) Data() interface{} {
	return self.data
}
