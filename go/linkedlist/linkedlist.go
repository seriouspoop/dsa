package linkedlist

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head *node
	size int
}

// Too much complex to debug and understand
func (n *node) Iter() func(func(*node) bool) {
	return func(yield func(*node) bool) {
		for n != nil {
			if !yield(n) {
				return
			}
			n = n.next
		}
	}
}

func NewLinkedList() *linkedList {
	return &linkedList{nil, 0}
}

func (l *linkedList) Populate(size int) {
	if l.head == nil && size > 0 {
		l.head = new(node)
	}
	head := l.head
	for i := range size {
		head.value = i + 1
		if i != size-1 {
			head.next = new(node)
			head = head.next
		}
	}
	l.size = size
}

func (l *linkedList) DisplayWithLoop() {
	//temporary iterator
	head := l.head
	for head != nil {
		fmt.Printf("%d->", head.value)
		head = head.next
	}
	fmt.Println(nil)
}

func (l *linkedList) DisplayWithIterator() {
	for node := range l.head.Iter() {
		fmt.Printf("%d->", node.value)
	}
	fmt.Println(nil)
}

func (l *linkedList) AppendNode(value int) {

	if l.head == nil {
		l.head = &node{value, nil}
		l.size++
		return
	}

	head := l.head
	// Iterate till the end of the list
	for head.next != nil {
		head = head.next
	}

	head.next = new(node)
	head.next.value = value
	l.size++
}

func (l *linkedList) InsertNode(value int) {
	// create new head
	node := &node{value, l.head}
	l.head = node
	l.size++
}

func (l *linkedList) DeleteNodeWithValue(value int) {
	var prev *node
	head := l.head

	for head.value != value {
		prev = head
		head = head.next
	}

	prev.next = head.next
	l.size--
}

func (l *linkedList) deleteNode(position int) {
	if position > l.size {
		fmt.Println("invalid position")
		return
	}

	var prev *node
	head := l.head

	// first position is head
	switch position {
	case 1:
		l.head = head.next
		l.size--
		return
	case -1:
		position = l.size
	}
	for range position - 1 {
		prev = head
		head = head.next
	}

	prev.next = head.next
	l.size--
}

func (l *linkedList) DeleteNodeAtBeginning() {
	l.deleteNode(1)
}

func (l *linkedList) DeleteNodeAtEnd() {
	l.deleteNode(-1)
}

func (l *linkedList) DeleteNodeAtPosition(pos int) {
	l.deleteNode(pos)
}

func (l *linkedList) InsertNodeAtPosition(value, pos int) {
	if pos > l.size {
		fmt.Println("invalid position")
		return
	}
	head := l.head
	// (pos - 1) for converting numbers to indexes, -1 for stopping at the previous node
	for range (pos - 1) - 1 {
		head = head.next
	}

	newNode := &node{value, head.next}
	head.next = newNode
	l.size++
}

func (l *linkedList) ReverseWithSlice() {
	// This initializes the slice with zero value. append won't work
	valueSlice := make([]int, l.size)

	head := l.head
	// append the first value
	for i := range l.size {
		valueSlice[i] = head.value
		head = head.next
	}

	head = l.head
	for i := range l.size {
		head.value = valueSlice[(l.size-1)-i]
		head = head.next
	}
}

func (l *linkedList) ReverseWithLinks() {
	var start *node = nil
	end := l.head

	for end != nil {
		temp := end
		end = end.next
		temp.next = start
		start = temp
	}

	l.head = start
}
