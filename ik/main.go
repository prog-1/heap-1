package main

import "fmt"

type Heap []int

func findMinIK(m int, d []int) (ikMin, days int) {
	var solve func(m, ikMin, days int, d []int) (int, int)
	solve = func(m, ikMin, days int, d []int) (int, int) {
		initial := d
		h := new(Heap)
		ik := ikMin
		for len(d) != 0 || len(*h) != 0 {
			if len(d) != 0 {
				var tmp []int
				if len(d) >= m {
					tmp = d[:m-len(*h)]
				} else {
					tmp = d
				}
				if len(d) >= m {
					d = d[m-len(*h):]
				} else {
					d = nil
				}
				for i := 0; i < len(tmp) && i < m; i++ {
					h.Push(tmp[i])
				}
			}
			check := len(*h)
			for i := 0; i < len(*h); i++ {
				x := h.Pop()
				if ik >= x {
					ik++
					i--
				} else {
					h.Push(x)
					break
				}
			}
			if len(*h) == check {
				return solve(m, ikMin+1, 0, initial)
			}
			days++
		}
		return ikMin, days
	}
	return solve(m, 1, 0, d)
}

func (h *Heap) Push(x int) {
	*h = append(*h, x)
	for i := len(*h) - 1; (*h)[i] < (*h)[(i-1)/2]; i = (i - 1) / 2 {
		(*h)[i], (*h)[(i-1)/2] = (*h)[(i-1)/2], (*h)[i]
	}
}

func (h *Heap) Pop() int {
	if len(*h) == 0 {
		return 0
	}
	x := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	var i, j int
	for i*2+1 < len(*h) {
		if i*2+2 == len(*h) {
			j = i*2 + 1
		} else {
			if (*h)[i*2+1] < (*h)[i*2+2] {
				j = i*2 + 1
			} else {
				j = i*2 + 2
			}
		}
		if (*h)[i] > (*h)[j] {
			(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
			i = j
		} else {
			break
		}
	}
	return x
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&d[i])
	}
	fmt.Println(findMinIK(m, d))
}
