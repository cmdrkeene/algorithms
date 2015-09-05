package algorithms

import (
	"reflect"
	"testing"
)

func TestLinkedList_Size(t *testing.T) {
	// given
	list := NewLinkedList()

	// when
	list.Add(1).Add(2).Add(3)

	// then
	Expect(t).Equal(list.Size(), 3)
}

func TestLinkedList_RemoveHead(t *testing.T) {
	// given
	list := NewLinkedList().Add(1).Add(2).Add(3)

	// when
	list.Remove(1)

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
	list.Remove(2)

	// then
	Expect(t).Equal(list.String(), "1->3")
}

func TestLinkedList_String(t *testing.T) {
	// given
	list := NewLinkedList()

	// when
	list.Add(1).Add(2).Add(3)

	// then
	Expect(t).Equal(list.String(), "1->2->3")
}

func Expect(t *testing.T) *expectation {
	return &expectation{t: t}
}

type expectation struct {
	t *testing.T
}

func (e *expectation) Equal(a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		e.t.Error(a, "does not equal", b)
	}
}
