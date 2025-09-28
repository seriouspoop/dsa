package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		name     string
		array    []int
		reversed []int
	}{
		{
			name:     "multiple elements",
			array:    []int{1, 23, 3, 4, 56},
			reversed: []int{56, 4, 3, 23, 1},
		},
		{
			name:     "single element",
			array:    []int{2},
			reversed: []int{2},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			arr := NewArray(test.array)

			arr.Reverse()
			assert.Equal(t, test.reversed, arr.arr)
		})
	}
}

func TestInsertManual(t *testing.T) {
	testCases := []struct {
		name     string
		value    int
		arr      []int
		new      []int
		position int
	}{
		{
			name:     "insert half",
			arr:      []int{1, 2, 3, 4, 5},
			new:      []int{1, 2, 9, 3, 4, 5},
			position: 5 / 2,
			value:    9,
		},
		{
			name:     "insert at 0",
			arr:      []int{1, 2, 3},
			new:      []int{0, 1, 2, 3},
			position: 0,
			value:    0,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			arr := NewArray(test.arr)
			arr.InsertManual(test.value, test.position)

			assert.Equal(t, test.new, arr.arr)
		})
	}
}

func TestInsertWithCopy(t *testing.T) {
	testCases := []struct {
		name     string
		value    int
		arr      []int
		new      []int
		position int
	}{
		{
			name:     "insert half",
			arr:      []int{1, 2, 3, 4, 5},
			new:      []int{1, 2, 9, 3, 4, 5},
			position: 5 / 2,
			value:    9,
		},
		{
			name:     "insert at 0",
			arr:      []int{1, 2, 3},
			new:      []int{0, 1, 2, 3},
			position: 0,
			value:    0,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			arr := NewArray(test.arr)
			arr.InsertWithCopy(test.value, test.position)

			assert.Equal(t, test.new, arr.arr)
		})
	}
}

func createTestArray(size int) *array {
	arr := &array{arr: make([]int, size)}
	for i := range size {
		arr.arr[i] = i
	}

	return arr
}

func BenchmarkInsertManual(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000, 100000, 1000000}

	for _, size := range sizes {
		position := size / 2
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				arr := createTestArray(size)
				b.StartTimer()
				arr.InsertManual(999, position)
			}
		})
	}
}

func BenchmarkInsertWithCopy(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000, 100000, 1000000}

	for _, size := range sizes {
		position := size / 2
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				arr := createTestArray(size)
				b.StartTimer()
				arr.InsertWithCopy(999, position)
			}
		})
	}
}
