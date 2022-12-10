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

func left(index int) int {
	return index*2 + 1
}

func right(index int) int {
	return index*2 + 2
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Pop() int {
	tmp := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	var current int
	for {
		left := left(current)
		if left+1 < len(*h) {
			l, r := (*h)[left], (*h)[left+1]
			if l < r {
				if l < (*h)[current] {
					(*h)[current], (*h)[left] = (*h)[left], (*h)[current]
					current = left + 1
				} else if r < (*h)[current] {
					(*h)[current], (*h)[left+1] = (*h)[left+1], (*h)[current]
					current = left + 2
				} else {
					break
				}
			} else {
				if r < (*h)[current] {
					(*h)[current], (*h)[left+1] = (*h)[left+1], (*h)[current]
					current = left + 2
				} else {
					break
				}

			}

		} else if left < len(*h) {
			if (*h)[current] > (*h)[left] {
				(*h)[current], (*h)[left] = (*h)[left], (*h)[current]
			}
			break
		} else {
			break
		}
	}
	return tmp
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
	fmt.Println(h.Pop())
	fmt.Println(h)
	fmt.Println(h.Pop())
	fmt.Println(h)
	fmt.Println(h.Pop())
	fmt.Println(h)
	fmt.Println(h.Pop())
	fmt.Println(h)
	fmt.Println(h.Pop())
	fmt.Println(h)
	fmt.Println(h.Pop())
	fmt.Println(h)
}
