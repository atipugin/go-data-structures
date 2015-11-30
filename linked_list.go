package ds

import (
	"errors"
)

type LinkedList struct {
	head *LinkedListNode
	size int
}

type LinkedListNode struct {
	key  int
	next *LinkedListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) append(key int) *LinkedListNode {
	node := &LinkedListNode{key: key}

	if list.head != nil {
		tail := list.head

		for tail.next != nil {
			tail = tail.next
		}

		tail.next = node
	} else {
		list.head = node
	}

	list.size = list.size + 1

	return node
}

func (list *LinkedList) prepend(key int) *LinkedListNode {
	node := &LinkedListNode{key: key}
	head := list.head
	list.head = node
	node.next = head
	list.size = list.size + 1

	return node
}

func (list *LinkedList) insertBefore(key int, before int) (*LinkedListNode, error) {
	var pred *LinkedListNode

	node := &LinkedListNode{key: key}
	next := list.head

	for next != nil && next.key != before {
		pred = next
		next = next.next
	}

	if next == nil {
		return nil, errors.New("Key not found")
	}

	node.next = next

	if pred != nil {
		pred.next = node
	} else {
		list.head = node
	}

	list.size = list.size + 1

	return node, nil
}

func (list *LinkedList) insertAfter(key int, after int) (*LinkedListNode, error) {
	node := &LinkedListNode{key: key}
	pred := list.head

	for pred != nil && pred.key != after {
		pred = pred.next
	}

	if pred == nil {
		return nil, errors.New("Key not found")
	}

	node.next = pred.next
	pred.next = node
	list.size = list.size + 1

	return node, nil
}

func (list *LinkedList) remove(key int) (*LinkedListNode, error) {
	var pred *LinkedListNode

	node := list.head

	for node != nil && node.key != key {
		pred = node
		node = node.next
	}

	if node == nil {
		return nil, errors.New("Key not found")
	}

	if pred != nil {
		pred.next = node.next
	} else {
		list.head = node.next
	}

	list.size = list.size - 1

	return node, nil
}

func (list *LinkedList) find(key int) (*LinkedListNode, error) {
	node := list.head

	for node != nil && node.key != key {
		node = node.next
	}

	if node == nil {
		return nil, errors.New("Key not found")
	}

	return node, nil
}
