package linkedlist

import (
	"errors"
)

// Element doc here
type Element struct {
	data int
	next *Element
}

// List doc here
type List struct {
	head *Element
	size int
}

// New doc here
func New(values []int) *List {
	list := List{}
	for _, v := range values {
		list.Push(v)
	}
	return &list
}

// Size doc here
func (list *List) Size() int {
	return list.size
}

// Push doc here
func (list *List) Push(value int) {
	defer func() { list.size++ }()

	ele := Element{data: value}

	if list.head == nil {
		list.head = &ele
		return
	}

	for e := list.head; ; e = e.next {
		if e.next == nil {
			e.next = &ele
			break
		}
	}
}

// Pop doc here
func (list *List) Pop() (value int, err error) {
	if list.head == nil {
		return 0, errors.New("")
	}
	defer func() { list.size-- }()

	if list.head.next == nil {
		value = list.head.data
		list.head = nil
		return value, nil
	}

	for e := list.head; ; e = e.next {
		if e.next.next == nil {
			value = e.next.data
			e.next = nil
			return value, nil
		}
	}
}

// Array doc here
func (list *List) Array() []int {
	array := make([]int, 0, list.size)
	for e := list.head; e != nil; e = e.next {
		array = append(array, e.data)
	}
	return array
}

// Reverse doc here
func (list *List) Reverse() *List {
	array := list.Array()
	for i := 0; i < len(array)/2; i++ {
		array[i], array[len(array)-1-i] = array[len(array)-1-i], array[i]
	}
	return New(array)
}
