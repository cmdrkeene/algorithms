package algorithms

import (
	"bytes"
	"fmt"
)

func NewLinkedList(elements ...interface{}) *LinkedListNode {
	head := &LinkedListNode{}
	return head.Append(elements...)
}

type LinkedListNode struct {
	data interface{}
	next *LinkedListNode
}

const (
	LessThan    = -1
	EqualTo     = 0
	GreaterThan = 1
)

// Compare returns -1, 0, 1 to mean less than, equal, greater than
// Cheat and cast to int
func (self *LinkedListNode) Compare(a, b interface{}) int {
	return compareInt(a.(int), b.(int))
}

func compareInt(a, b int) int {
	if a < b {
		return LessThan
	}

	if a > b {
		return GreaterThan
	}

	return EqualTo
}

// Partition returns a new LinkedList such that: [ before < pivot < after ]
// This doesn't sort, it just buckets according to less / greater than
func (self *LinkedListNode) Partition(pivot interface{}) *LinkedListNode {
	// Not found
	if self.Find(pivot) == nil {
		return self
	}

	// Two buckets, ignoring found
	before := NewLinkedList()
	after := NewLinkedList()
	node := self
	for {
		// end of list
		if node == nil {
			break
		}

		// put into appropriate bucket
		switch self.Compare(node.data, pivot) {
		case EqualTo:
			before.Append(node.data)
		case LessThan:
			before = before.Prepend(node.data)
		case GreaterThan:
			after = after.Append(node.data)
		}

		// advance cusror
		node = node.next
	}

	// merge
	before.Tail().next = after
	return before
}

func (self *LinkedListNode) Prepend(data interface{}) *LinkedListNode {
	if self.data == nil {
		self.data = data
		return self
	}
	return &LinkedListNode{data: data, next: self}
}

// Find returns first node, if any, that contains data
func (self *LinkedListNode) Find(data interface{}) *LinkedListNode {
	node := self
	for {
		if node == nil {
			return nil
		}
		if node.data == data {
			return node
		}
		node = node.next
	}
}

// FindLoop returns the node, if any, where the List starts looping
func (self *LinkedListNode) FindLoop() *LinkedListNode {
	// Use the runner method to look ahead into the list
	slow := self
	fast := self

	// first pass
	for {
		// no loop exists if end is reached
		if slow.next == nil || fast.next == nil {
			return nil
		}

		slow = slow.next      // one element at a time
		fast = fast.next.next // two elements at a time

		// collision
		if slow == fast {
			break
		}
	}

	// second pass starts slow at head, advances both at same speed until collision
	slow = self
	for {
		slow = slow.next
		fast = fast.next

		// collision
		if slow == fast {
			break
		}
	}

	// advance slow and fast at same speed until collision
	// at collision, slow is now at loop start
	return slow
}

func (self *LinkedListNode) Append(elements ...interface{}) *LinkedListNode {
	for _, e := range elements {
		// head is empty
		if self.data == nil {
			self.data = e
		} else {
			self.Tail().next = NewLinkedList(e)
		}
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

		if node.data == nil {
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
