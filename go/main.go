package main

import (
	"github.com/seriouspoop/dsa/arrays"
	"github.com/seriouspoop/dsa/linkedlist"
)

func main() {
	// Arrays
	arr := arrays.NewArray([]int{1, 2, 3, 4, 5})

	arr.Display()
	arr.Reverse()
	arr.Display()

	arr.InsertManual(99, 2)
	arr.Display()

	arr.InsertWithCopy(88, 3)
	arr.Display()

	length := arr.Len()
	println(length)

	// Linked List
	ll := linkedlist.NewLinkedList()
	ll.Populate(5)
	ll.DisplayWithLoop()
	ll.AppendNode(99)
	ll.DisplayWithIterator()
	ll.InsertNode(0)
	ll.DisplayWithIterator()
	ll.DeleteNodeWithValue(99)
	ll.DisplayWithIterator()
	ll.DeleteAtBeginning()
	ll.DisplayWithIterator()
	ll.DeleteAtEnd()
	ll.DisplayWithIterator()
}
