package arrays

import "fmt"

// Fundamental operations on arrays
//
// Operations covered:
//  1. reverse
//  1. insertion
//  1. deletion
type array struct {
	arr []int
}

func NewArray(arr []int) *array {
	return &array{
		arr: arr,
	}
}

func (a *array) Len() int {
	return len(a.arr)
}

func (a *array) Display() {
	fmt.Println(a.arr)
}

// Use 2 pointers i, j. One traversing from the beginning and other from the end.
// Cut the array length in half and swap the 2 end marked by the pointers.
// Gradually move inwards.
func (a *array) Reverse() {
	for i := 0; i < len(a.arr)/2; i++ {
		j := len(a.arr) - 1 - i
		a.arr[i], a.arr[j] = a.arr[j], a.arr[i]
	}
}

// Insert traverses the array from the end and slides each element one space father
// making space for the new element at the position.
// Then the new element is inserted at the given position.
func (a *array) InsertManual(value int, position int) {
	if position < 0 || position > len(a.arr) {
		panic("Error index out of bounds")
	}

	// increase the length of the array
	a.arr = append(a.arr, 0)

	// shift elements right to make gap
	for i := len(a.arr) - 1; i > position; i-- {
		a.arr[i] = a.arr[i-1]
	}
	a.arr[position] = value
}

func (a *array) InsertWithCopy(value int, position int) {
	if position < 0 || position > len(a.arr) {
		panic("Error index out of bounds")
	}

	// increase the length of the slice/array
	a.arr = append(a.arr, 0)

	// copy the slice after the position and shift it one position ahead
	copy(a.arr[position+1:], a.arr[position:])

	// insert the value at the position
	a.arr[position] = value
}
