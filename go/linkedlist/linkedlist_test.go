package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func llToArray(list *linkedList) []int {
	lArray := make([]int, list.size)

	head := list.head
	for i := range list.size {
		lArray[i] = head.value
		head = head.next
	}

	return lArray
}

func TestPopulate(t *testing.T) {
	testCases := []struct {
		name string
		size int
		list []int
	}{
		{
			name: "5 elements",
			size: 5,
			list: []int{1, 2, 3, 4, 5},
		},
		{
			name: "2 elements",
			size: 2,
			list: []int{1, 2},
		},
		{
			name: "0 elements, empty erray",
			size: 0,
			list: []int{},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			ll := NewLinkedList()
			ll.Populate(test.size)

			assert.ElementsMatch(t, test.list, llToArray(ll))
		})
	}
}

func TestAppendNode(t *testing.T) {
	testCases := []struct {
		name    string
		size    int
		elem    int
		newSize int
		list    []int
	}{
		{
			name:    "5 elements, append 6",
			size:    5,
			elem:    6,
			newSize: 6,
			list:    []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:    "0 elements, append 10",
			size:    0,
			newSize: 1,
			list:    []int{10},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			ll := NewLinkedList()
			ll.Populate(test.size)
			ll.AppendNode(test.elem)

			assert.ElementsMatch(t, test.list, llToArray(ll))
			assert.Equal(t, test.newSize, ll.size)
		})
	}
}

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
