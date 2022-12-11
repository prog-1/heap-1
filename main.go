package main

import (
	"fmt"
	"math"
)

type Heap []int

func main() {
	fmt.Println(IngusCoefficient([]int{3, 1, 7, 2, 6, 3}, 3)) // 2 3
	//fmt.Println(IngusCoefficient([]int{7, 6, 5, 4, 3}, 9)) // 3 1
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
	if len(*h) == 0 {
		panic("Popping element from an empty Heap")
	}
	output := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	if len(*h) == 0 {
		return output
	}

	i := 0
	for {
		l, r := math.MaxInt, math.MaxInt
		if (2*i + 1) < len(*h) { // left child exists
			l = (*h)[2*i+1]
		}
		if (2*i + 2) < len(*h) { // right child exists
			r = (*h)[2*i+2]
		}
		if (*h)[i] < l && (*h)[i] < r {
			break
		}
		p := i
		if l <= r {
			i = 2*i + 1
		} else {
			i = 2*i + 2
		}
		(*h)[p], (*h)[i] = (*h)[i], (*h)[p]
	}

	return output
}

func IngusCoefficient(tasks []int, m int) (minIC, days int) {
	var h Heap
	h.Push(tasks[0])
	curIC := 0
	for i := 1; len(h) != 0; {
		for ; len(h) < m && i < len(tasks); i++ {
			h.Push(tasks[i])
		}
		for min := h.Pop(); len(h) > 0; min = h.Pop() {
			if min <= curIC { // can solve
				curIC++
			} else { // can't solve
				minIC = minIC + (min - curIC)
				curIC = min + 1
			}
		}
	}
	return minIC, days
}
