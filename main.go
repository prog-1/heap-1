package main

import (
	"bufio"
	"fmt"
	"os"
)

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
	i := 0
	if len(*h) == 0 {
		return 0
	}

	(*h)[i], (*h)[len(*h)-1] = (*h)[len(*h)-1], (*h)[i]
	*h = (*h)[:len(*h)-1]

	for {
		l, r := 2*i, 2*i+1
		min := i
		if l < len(*h) && (*h)[i] > (*h)[l] {
			min = l
		}
		if r < len(*h) && (*h)[l] > (*h)[r] {
			min = r
		}
		if min != i {
			(*h)[i], (*h)[min] = (*h)[min], (*h)[i]
			i = min
			continue

		}
		break
	}
	if len(*h) == 0 {
		return 0
	}
	return (*h)[0]
}

func MinIK(uzd []int, m int) int {
	var queue Heap
	for i := 0; i < m && len(uzd) != 0; i++ {
		queue.Push(uzd[0])
		uzd = uzd[1:]
	}
	minIk := 1
	ik := minIk
	for len(uzd) > 0 || len(queue) != 0 {
		min := queue[0]
		if min <= ik {
			_ = queue.Pop()
			ik++

			if len(uzd) > 0 {
				queue.Push(uzd[0])
				uzd = uzd[1:]
			}

			continue
		}
		minIk += (min - ik)
		ik += (min - ik)

	}
	return minIk
}

func MinDays(uzd []int, m, ik int) {
	var queue Heap
	ikk := ik
	days := 0
	for len(uzd) > 0 || len(queue) > 0 {
		days++
		for len(queue) < m && len(uzd) > 0 {
			queue.Push(uzd[0])
			uzd = uzd[1:]
		}
		// fmt.Println(len(uzd), len(queue))
		for i := m; i > 0 && len(queue) > 0; i-- {
			min := queue[0]
			// fmt.Println(min, ik, "             ", days)
			if min <= ik {
				_ = queue.Pop()
				ik++
				continue
			}
			break
		}
	}

	fmt.Println(ikk, days)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n, &m)
	uzd := make([]int, n)
	for i := range uzd {
		fmt.Fscan(r, &uzd[i])
	}
	// fmt.Println(uzd)
	ik := MinIK(uzd, m)

	MinDays(uzd, m, ik)
}

// func main() {
// 	var h Heap

// 	h.Push(1)
// 	h.Push(3)
// 	h.Push(7)
// 	h.Push(5)
// 	h.Push(9)
// 	// h.Push(4)
// 	fmt.Println(h)

// }
