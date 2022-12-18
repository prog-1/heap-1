package main

import "fmt"

type Heap []int

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Push(x int) {
	node := len(*h)
	*h = append(*h, x)
	parent := (node - 1) / 2
	for (*h)[parent] > (*h)[node] {
		(*h)[parent], (*h)[node] = (*h)[node], (*h)[parent]
		parent = (node - 1) / 2
	}
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Pop() int {
	min := (*h)[0]
	node := len(*h)
	(*h)[0] = (*h)[node-1]
	h.delMin(0)
	(*h) = (*h)[:node-1]
	return min
}

func (h *Heap) delMin(node int) {
	var min int
	length := len(*h)
	left, right := 2*node+1, 2*node+2

	if right >= length && left >= length {
		return
	} else if right >= length && left < length {
		min = left
	} else {
		if (*h)[left] < (*h)[right] {
			min = left
		} else {
			min = right
		}
	}
	if (*h)[min] < (*h)[node] {
		(*h)[min], (*h)[node] = (*h)[node], (*h)[min]
		h.delMin(min)
	}
}

func main() {
	var h Heap
	h.Push(1)
	h.Push(3)
	h.Push(3)
	h.Push(9)
	h.Push(7)
	h.Push(5)
	h.Push(4)
	fmt.Println(h.Pop())
	fmt.Println(h)
}
