package algorithms

import "testing"

func TestQueue(t *testing.T) {
	queue := NewQueue()

	// Enqueue
	queue.Enqueue(1, 2, 3)

	// Size
	Expect(t).Equal(queue.Size(), 3)

	// Peek

	// Dequeue
	Expect(t).Equal(queue.Dequeue(), 1)
	Expect(t).Equal(queue.Size(), 2)

	// Dequeue until empty
	Expect(t).Equal(queue.Dequeue(), 2)
	Expect(t).Equal(queue.Size(), 1)
	Expect(t).Equal(queue.Dequeue(), 3)
	Expect(t).Equal(queue.Size(), 0)

	// Empty
	Expect(t).Equal(queue.Empty(), true)

	// Dequeue when empty
	Expect(t).Equal(queue.Dequeue(), nil)
}
