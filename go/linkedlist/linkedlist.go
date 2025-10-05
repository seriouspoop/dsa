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
	return &linkedList{new(node), 0}
}

func (l *linkedList) Populate(size int) {
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
		if head == nil {
			fmt.Println("invalid position")
			return
		}

		prev = head
		head = head.next
	}

	prev.next = head.next
	l.size--
}

func (l *linkedList) DeleteAtBeginning() {
	l.deleteNode(1)
}

func (l *linkedList) DeleteAtEnd() {
	l.deleteNode(-1)
}
