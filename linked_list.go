package algorithms

import (
	"bytes"
	"fmt"
)

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

type LinkedList struct {
	head *LinkedListNode
}

func (self *LinkedList) Add(data interface{}) *LinkedList {
	tail := &LinkedListNode{data: data, next: nil}

	// base case
	if self.head == nil {
		self.head = tail
		return self
	}

	// n-th case
	node := self.head
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
func (self *LinkedList) Remove(data interface{}) *LinkedList {
	// List is empty
	if self.head == nil {
		return self
	}

	// List head removed
	if self.head.data == data {
		self.head = self.head.next
		return self
	}

	// Scan list body
	node := self.head
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

func (self *LinkedList) Size() int {
	if self.head == nil {
		return 0
	}

	var size int
	node := self.head
	for {
		if node == nil {
			break
		}
		size++
		node = node.next
	}

	return size
}

func (self *LinkedList) String() string {
	buf := bytes.NewBufferString("")

	node := self.head
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

type LinkedListNode struct {
	data interface{}
	next *LinkedListNode
}
