package linkedlist

import "fmt"

type node struct {
	value int
	next  *node
}

type linkedList struct {
	list *node
}

// Too much complex to debug and understand
func (n *node) Iter() func(func(*node) bool) {
	return func(yield func(*node) bool) {
		for n != nil {
			if !yield(n) {
				return
			}
			n.value += 10
			n = n.next
		}
	}
}

func NewLinkedList(value int) *linkedList {
	return &linkedList{
		list: &node{
			value,
			&node{
				value + 10,
				nil,
			},
		},
	}
}

func (l *linkedList) Head() *node {
	return l.list
}

func (l *linkedList) Display() {
	// for node := range l.list.Iter() {
	// 	fmt.Printf("%d->", node.value)
	// }

	//temporary iterator
	head := l.list
	for head != nil {
		fmt.Printf("%d->", head.value)
		head = head.next
	}
	fmt.Println(nil)
}
