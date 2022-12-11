package main

import "fmt"

type Heap []int

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Push(x int) {
	i := len(*h)
	*h = append(*h, x)
	if i == 0 {
		return
	}
	for (*h)[i] < (*h)[(i-1)/2] {
		(*h)[i], (*h)[(i-1)/2] = (*h)[(i-1)/2], (*h)[i]
		i = (i - 1) / 2
	}

}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Pop() int {
	l := len(*h)
	if l == 0 {
		return 0
	}
	min := (*h)[0]
	(*h)[0] = (*h)[l-1]
	*h = (*h)[:l-1]
	i := 0
	for h.isvalid(2*i+1) && (*h)[i] > (*h)[2*i+1] || h.isvalid(2*i+2) && (*h)[i] > (*h)[2*i+2] {
		if h.isvalid(2*i+1) && !h.isvalid(2*i+2) && (*h)[2*i+1] < (*h)[i] {
			(*h)[i], (*h)[2*i+1] = (*h)[2*i+1], (*h)[i]
			i = (i * 2) + 1
		} else if !h.isvalid(2*i+1) && h.isvalid(2*i+2) && (*h)[2*i+2] < (*h)[i] {
			(*h)[i], (*h)[2*i+2] = (*h)[2*i+2], (*h)[i]
			i = (i * 2) + 2
		} else if (*h)[2*i+1] < (*h)[2*i+2] {
			(*h)[i], (*h)[2*i+1] = (*h)[2*i+1], (*h)[i]
			i = (i * 2) + 1
		} else if (*h)[2*i+1] > (*h)[2*i+2] {
			(*h)[i], (*h)[2*i+2] = (*h)[2*i+2], (*h)[i]
			i = (i * 2) + 2
		}
	}
	return min
}
func (h *Heap) isvalid(el int) bool {
	return el >= 0 && el < len((*h))
}
func main() {
	var h Heap
	h.Push(1)
	h.Push(2)
	h.Push(3)
	// h.Push(7)
	// h.Push(5)
	// h.Push(4)
	fmt.Println(h)
	fmt.Println(h.Pop())
	fmt.Println(h)
}
