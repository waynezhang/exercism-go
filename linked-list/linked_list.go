package linkedlist

import (
	"errors"
)

// ErrEmptyList doc here
var ErrEmptyList = errors.New("")

// Node doc here
type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

// Next doc here
func (node *Node) Next() *Node {
	return node.next
}

// Prev doc here
func (node *Node) Prev() *Node {
	return node.prev
}

// List doc here
type List struct {
	head *Node
	tail *Node
	size int
}

// PushFront doc here
func (list *List) PushFront(data interface{}) {
	node := Node{Val: data, next: list.head}

	if list.size == 0 {
		list.head = &node
		list.tail = &node
	} else {
		list.head.prev = &node
		list.head = &node
	}
	list.size++
}

// PushBack doc here
func (list *List) PushBack(data interface{}) {
	node := Node{Val: data, prev: list.tail}

	if list.size == 0 {
		list.head = &node
		list.tail = &node
	} else {
		list.tail.next = &node
		list.tail = &node
	}
	list.size++
}

// PopFront doc here
func (list *List) PopFront() (data interface{}, err error) {
	if list.head == nil {
		return 0, ErrEmptyList
	}

	data = list.head.Val

	list.size--
	if list.size == 0 {
		list.head, list.tail = nil, nil
	} else {
		list.head = list.head.next
		list.head.prev = nil
	}
	return
}

// PopBack doc here
func (list *List) PopBack() (data interface{}, err error) {
	if list.tail == nil {
		return 0, ErrEmptyList
	}

	data = list.tail.Val

	list.size--
	if list.size == 0 {
		list.head, list.tail = nil, nil
	} else {
		list.tail = list.tail.prev
		list.tail.next = nil
	}

	return
}

// First doc here
func (list *List) First() *Node {
	return list.head
}

// Last doc here
func (list *List) Last() *Node {
	return list.tail
}

// Reverse doc here
func (list *List) Reverse() {
	for e := list.head; e != nil; e = e.prev {
		e.next, e.prev = e.prev, e.next
	}
	list.head, list.tail = list.tail, list.head
}

// NewList doc here
func NewList(datas ...interface{}) *List {
	list := &List{}
	for _, data := range datas {
		list.PushBack(data)
	}
	return list
}
