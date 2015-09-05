package algorithms

import "testing"

func TestStack(t *testing.T) {
	// New
	stack := NewStack()

	// Size
	Expect(t).Equal(stack.Size(), 0)

	// Push
	stack.Push(1, 2, 3)

	// Size
	Expect(t).Equal(stack.Size(), 3)

	// Pop
	Expect(t).Equal(stack.Pop(), 3)
	Expect(t).Equal(stack.Size(), 2)

	// Peek
	Expect(t).Equal(stack.Peek(), 2)

	// Pop until empty
	stack.Pop()
	stack.Pop()
	Expect(t).Equal(stack.Pop(), nil)
	Expect(t).Equal(stack.Size(), 0)
}
