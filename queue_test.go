package algorithms

import "testing"

func TestQueue_Stack(t *testing.T) {
	runQueueTest(t, NewStackQueue())
}

func TestQueue_LinkedList(t *testing.T) {
	runQueueTest(t, NewLinkedListQueue())
}

func runQueueTest(t *testing.T, queue Queue) {
	// Empty
	Expect(t).Equal(queue.Size(), 0)

	// Enqueue
	queue.Enqueue(1, 2, 3)

	// Size
	Expect(t).Equal(queue.Size(), 3)

	// Peek
	Expect(t).Equal(queue.Peek(), 1)

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

	// Enqueue after empty
	queue.Enqueue(1)
	Expect(t).Equal(queue.Size(), 1)
}
