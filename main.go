package main

import (
	"fmt"
	"math"
)

type Heap []int

func main() {
	fmt.Println(IngusCoefficient(processInput()))
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

func processInput() (tasks []int, m int) {
	var taskCount int
	fmt.Scanln(&taskCount, &m)
	tasks = make([]int, taskCount)
	for i := range tasks {
		fmt.Scan(&tasks[i])
	}
	return tasks, m
}

func IngusCoefficient(tasks []int, m int) (minIC, days int) {
	var h Heap
	for i, curIC := 0, 0; len(h) != 0 || i < len(tasks); {
		for ; len(h) < m && i < len(tasks); i++ {
			h.Push(tasks[i])
		}
		prev := len(h)
		for ; len(h) > 0 && h[0] <= curIC; curIC++ {
			h.Pop()
		}
		if len(h) == prev { // solved nothing
			minIC = minIC + (h[0] - curIC)
			curIC = h[0]
		}
	}

	for i, curIC := 0, minIC; len(h) != 0 || i < len(tasks); {
		for ; len(h) < m && i < len(tasks); i++ {
			h.Push(tasks[i])
		}
		for ; len(h) > 0 && h[0] <= curIC; curIC++ {
			h.Pop()
		}
		days++
	}

	return minIC, days
}
