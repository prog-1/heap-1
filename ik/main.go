package main

import "fmt"

type Heap []int

func (h *Heap) Push(x int) {
	node := len(*h)
	*h = append(*h, x)
	parent := (node - 1) / 2
	for (*h)[parent] > (*h)[node] {
		(*h)[parent], (*h)[node] = (*h)[node], (*h)[parent]
		parent = (node - 1) / 2
	}
}

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

func findIK(n, m int, s []int) int {
	var ik, res int
	var tree Heap
	tree = append(tree, s[0])
	for i := 1; i < m; i++ {
		tree.Push(s[i])
	}
	for len(tree) > 0 {
		var i int
		minVal := tree.Pop()
		d := minVal - ik
		if d > 0 {
			ik += d
			res++
			if i < n {
				tree.Push(s[i])
				i++
			}
		}
		ik += minVal
	}
	return res
}

func main() {

	min := findIK(6, 3, []int{3, 1, 7, 2, 6, 3})
	fmt.Println(min)
}
