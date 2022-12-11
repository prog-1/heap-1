package main

import "fmt"

type Heap []int

func (h *Heap) Push(x int) {
	i := len(*h)
	(*h) = append((*h), x)
	if i == 0 {
		return
	}
	for (*h)[(i-1)/2] > (*h)[i] {
		(*h)[(i-1)/2], (*h)[i] = (*h)[i], (*h)[(i-1)/2]
		i = (i - 1) / 2
	}
}

func (h *Heap) Pop() int {
	if len(*h) == 0 {
		return 0
	}
	if len(*h) == 1 {
		return (*h)[0]
	}
	i, min := 1, (*h)[0]
	(*h)[0], (*h) = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	for (*h)[i-1] > (*h)[i*2-1] || (*h)[i-1] > (*h)[i*2] {
		if (*h)[i*2] > (*h)[i*2-1] {
			(*h)[i-1], (*h)[i*2-1] = (*h)[i*2-1], (*h)[i-1]
			i = i * 2
		} else {
			(*h)[i-1], (*h)[i*2] = (*h)[i*2], (*h)[i-1]
			i = i*2 + 1
		}
	}
	return min
}

func main() {
	var h Heap
	h.Push(1)
	h.Push(3)
	h.Push(9)
	h.Push(7)
	h.Push(5)
	h.Push(10)
	h.Push(2)
	h.Push(4)
	h.Push(70)
	h.Push(11)
	h.Push(8)
	fmt.Println(h)
	fmt.Println(h.Pop())
	fmt.Println(h)
}
