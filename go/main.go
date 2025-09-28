package main

import (
	"github.com/seriouspoop/dsa/arrays"
)

func main() {
	arr := arrays.NewArray([]int{1, 2, 3, 4, 5, 6})

	arr.Display()
	arr.Reverse()
	arr.Display()
	arr.InsertManual(9, arr.Len()/2)
	arr.Display()
}
