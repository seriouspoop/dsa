package linkedlist

import (
	"fmt"
	"testing"
)

func createLL(size int) *linkedList {
	ll := new(linkedList)
	ll.head = new(node)

	head := ll.head
	for i := range size {
		head.value = i
		if i != size-1 {
			head.next = new(node)
			head = head.next
		}
	}

	return ll
}

func BenchmarkDisplayWithLoop(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000, 100000, 1000000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				ll := createLL(size)
				b.StartTimer()

				ll.DisplayWithLoop()
			}
		})
	}
}

func BenchmarkDisplayWithIterator(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000, 100000, 1000000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				ll := createLL(size)
				b.StartTimer()

				ll.DisplayWithIterator()
			}
		})
	}
}
