package main

import (
	"fmt"
)

type Heap []int

func main() {
	h := Heap([]int{1, 3, 4, 7, 5, 9})
	//h.Pop()
	fmt.Println([]int(h))
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Push(x int) {
	// Add new element to the end of the Heap
	// compare new element's value with its parent
	// if new element < parent -> swap them
	*h = Heap(append(*h, x))
	for i := len(*h) - 1; i > 0 && (*h)[i] < (*h)[(i-1)/2]; i = (i - 1) / 2 {
		(*h)[i], (*h)[(i-1)/2] = (*h)[(i-1)/2], (*h)[i]
	}
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Pop() int {
	// Save the root and place last element at its place
	// If root is smaller than its children -> return
	// Pick the smallest child and swap it with the root
	// Do it until we are not a leaf
	if (*h) == nil {
		return 0
	}
	output := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	for i := 0; ; {
		p := (*h)[i]
		if (2*i + 2) > len(*h) { // No right child
			if (2*i+1) > len(*h) || p <= (*h)[2*i+1] { // No left child or parent is greater
				break
			}
			(*h)[i], (*h)[2*i+1] = (*h)[2*i+1], (*h)[i]
			i = 2*i + 1
		} else { // Has both children
			l := (*h)[2*i+1]
			r := (*h)[2*i+2]
			if p < l && p < r {
				break
			}
			if l < r {
				(*h)[i], (*h)[2*i+1] = (*h)[2*i+1], (*h)[i]
				i = 2*i + 1
			} else {
				(*h)[i], (*h)[2*i+2] = (*h)[2*i+2], (*h)[i]
				i = 2*i + 2
			}
		}
	}
	return output
}
