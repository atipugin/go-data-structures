package ds

import (
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	if NewLinkedList() == nil {
		t.Fail()
	}
}

func TestAppend(t *testing.T) {
	list := NewLinkedList()
	list.append(1)
	if list.head.key != 1 {
		t.Fail()
	}
}

func TestPrepend(t *testing.T) {
	list := NewLinkedList()
	list.append(2)
	list.prepend(1)
	if list.head.key != 1 {
		t.Fail()
	}
}

func TestInsertBefore(t *testing.T) {
	list := NewLinkedList()
	list.append(2)
	list.insertBefore(1, 2)
	if list.head.key != 1 {
		t.Fail()
	}

	_, err := list.insertBefore(4, 3)
	if err == nil {
		t.Fail()
	}
}

func TestInsertAfter(t *testing.T) {
	list := NewLinkedList()
	list.append(1)
	list.insertAfter(2, 1)
	if list.head.key != 1 {
		t.Fail()
	}

	_, err := list.insertAfter(3, 4)
	if err == nil {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	list := NewLinkedList()
	list.append(1)
	node, _ := list.remove(1)
	if node.key != 1 {
		t.Fail()
	}

	_, err := list.remove(2)
	if err == nil {
		t.Fail()
	}
}

func TestFind(t *testing.T) {
	list := NewLinkedList()
	list.append(1)
	node, _ := list.find(1)
	if node.key != 1 {
		t.Fail()
	}

	_, err := list.find(2)
	if err == nil {
		t.Fail()
	}
}
