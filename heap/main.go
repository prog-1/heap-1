package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Heap []int

func Parent(i int) int {
	return (i - 1) / 2
}

func L(i int) int {
	return 2*i + 1
}
func R(i int) int {
	return 2*i + 2
}

func (h *Heap) Push(x int) {
	i := len(*h) - 1
	*h = append(*h, x)
	for i > 0 && ((*h)[Parent(i)] > (*h)[i]) {
		(*h)[Parent(i)], (*h)[i] = (*h)[i], (*h)[Parent(i)]
		i = Parent(i)
	}
}

func (h *Heap) Pop() int {
	min := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	if len(*h) > 0 {
		siftDown(0, h)
	}
	return min
}

func siftDown(i int, h *Heap) {
	var min int
	if R(i) >= len(*h) {
		if L(i) >= len(*h) {
			return
		}
		min = L(i)
	} else {
		//if (*h)[L(i)] <= (*h)[R(i)] {
		//	min = L(i)
		//}
		min = R(i)
	}
	if (*h)[i] > (*h)[min] {
		(*h)[i], (*h)[min] = (*h)[min], (*h)[i]
		siftDown(min, h)
	}
}

func ik(r io.Reader) (int, int) {
	var n, m, days, minik int
	fmt.Fscan(r, &n, &m)
	var allEx []int
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &allEx[i])
	}
	var dailyEx Heap

	if m > len(allEx) {
		m = len(allEx)
	}

	for _, i := range allEx {
		if len(dailyEx) < m {
			dailyEx.Push(allEx[i])
		}
		if len(dailyEx) == m {
			days++
			if minik <= dailyEx[0] {
				minik = dailyEx[0]
			} else {
				minik++
				dailyEx.Pop()
			}
		}
	}
	return days, minik
}

func main() {
	var h Heap
	var z Heap
	h = append(h, 1, 5, 3, 8, 7, 4, 9, 13, 15)
	h.Push(16)
	z.Push(2)
	z.Push(7)
	fmt.Println(h, z)
	minh := h.Pop()
	fmt.Println(minh)
	fmt.Println(h)
	minz := z.Pop()
	fmt.Println(minz)
	fmt.Println(z)
	r := bufio.NewReader(os.Stdout)
	d, ik := ik(r)
	fmt.Println(d, ik)
}
