package main

import "fmt"

type Heap []int

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Push(x int) {
	*h = append(*h, x)
	current := len(*h) - 1
	for p := parent(current); (*h)[p] > x; p, current = parent(p), p {
		(*h)[current], (*h)[p] = (*h)[p], (*h)[current]
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Pop() int {
	return 0
}

func main() {
	var h Heap
	h.Push(1)
	h.Push(3)
	h.Push(9)
	h.Push(7)
	h.Push(5)
	h.Push(-1)

	fmt.Println(h)
}
