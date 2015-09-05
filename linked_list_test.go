package algorithms

import "testing"

func TestLinkedList_FindLoopWithoutLoop(t *testing.T) {
	if NewLinkedList(1, 2, 3).FindLoop() != nil {
		t.Error("no loop exists")
	}
}

func TestLinkedList_FindLoop(t *testing.T) {
	// given
	list := NewLinkedList(1, 2, 3, 4, 5)

	// when
	list.Tail().next = list.Find(3)

	// then
	Expect(t).Equal(list.FindLoop(), list.Find(3))
}

func TestLinkedList_Head(t *testing.T) {
	Expect(t).Equal(
		NewLinkedList(1, 2, 3).Head().Data().(int),
		1,
	)
}

func TestLinkedList_Tail(t *testing.T) {
	Expect(t).Equal(
		NewLinkedList(1, 2, 3).Tail().Data().(int),
		3,
	)
}

func TestLinkedList_Size(t *testing.T) {
	// given
	list := NewLinkedList(1, 2, 3)

	// when
	size := list.Size()

	// then
	Expect(t).Equal(size, 3)
}

func TestLinkedList_RemoveHead(t *testing.T) {
	// given
	list := NewLinkedList().Add(1).Add(2).Add(3)

	// when
	list = list.Remove(1)

	// then
	Expect(t).Equal(list.String(), "2->3")
}

func TestLinkedList_RemoveTail(t *testing.T) {
	// given
	list := NewLinkedList().Add(1).Add(2).Add(3)

	// when
	list.Remove(3)

	// then
	Expect(t).Equal(list.String(), "1->2")
}

func TestLinkedList_RemoveNotFound(t *testing.T) {
	// given
	list := NewLinkedList().Add(1).Add(2).Add(3)

	// when
	list.Remove(4)

	// then
	Expect(t).Equal(list.String(), "1->2->3")
}

func TestLinkedList_RemoveBody(t *testing.T) {
	// given
	list := NewLinkedList().Add(1).Add(2).Add(3)

	// when
	list = list.Remove(2)

	// then
	Expect(t).Equal(list.String(), "1->3")
}

func TestLinkedList_String(t *testing.T) {
	Expect(t).Equal(
		NewLinkedList(1, 2, 3).String(),
		"1->2->3",
	)
}
