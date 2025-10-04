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
	ll := linkedlist.NewLinkedList(10)
	ll.Display()
}
