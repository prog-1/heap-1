package main

import "fmt"

type Heap []int

func main() {
	h := Heap([]int{1, 3, 9, 7, 5})
	h.Push(4)
	fmt.Println([]int(h))
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Push(x int) {
	// Add new element to the end of the Heap
	// compare new element's value with its parent
	// if new element < parent -> swap them
	*h = Heap(append([]int(*h), x))
	i := len([]int(*h)) - 1
	for ; i > 0 && []int(*h)[i] < []int(*h)[(i-1)/2]; i = (i - 1) / 2 {
		[]int(*h)[i], []int(*h)[(i-1)/2] = []int(*h)[(i-1)/2], []int(*h)[i]
	}
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Pop() int {

}
