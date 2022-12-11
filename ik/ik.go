package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		i := (*h)[0]
		(*h) = (*h)[:len(*h)-1]
		return i
	}
	i, min := 0, (*h)[0]
	(*h)[0], (*h) = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	for len(*h)-1 >= i*2+1 {
		if len(*h)-1 == i*2+1 && (*h)[i] > (*h)[i*2+1] {
			(*h)[i], (*h)[i*2+1] = (*h)[i*2+1], (*h)[i]
			return min
		} else if len(*h)-1 == i*2+1 {
			return min
		} else if ((*h)[i] > (*h)[i*2+1] || (*h)[i] > (*h)[i*2+2]) && (*h)[i*2+2] > (*h)[i*2+1] {
			(*h)[i], (*h)[i*2+1] = (*h)[i*2+1], (*h)[i]
			i = i*2 + 1
		} else {
			(*h)[i], (*h)[i*2+2] = (*h)[i*2+2], (*h)[i]
			i = i*2 + 2
		}
	}
	return min
}

func (h *Heap) fill(i, experday int, ex []int) int {
	fmt.Println(h)
	fmt.Println(len(*h), experday, len(ex)-1, i)
	for len(*h) < experday && len(ex)-1 != i {
		i++
		fmt.Println(i, ex, ex[i])
		h.Push(ex[i])
	}
	if len(ex)-1 == i && len(*h) < experday {
		h.Push(ex[i])
	}
	fmt.Println(h, "**")
	return i
}

func (h *Heap) findMinIK(experday, i int, ex []int) (int, int) {
	ik, minik, dayneed := 1, 1, 0
	for {
		if len(ex)-1 == i && len(*h) == 0 {
			return minik, dayneed
		}
		if len(*h) == 0 {
			i = h.fill(i, experday, ex)
			dayneed++
		}
		fmt.Println(ik, "en")
		curr := h.Pop()
		fmt.Println("need", curr)
		if curr > ik && (len(*h) == experday || len(ex)-1 == i) {
			minik++
			*h, i, ik, dayneed = nil, 0, minik, 0
			fmt.Println(minik, "!")
		} else if curr > ik {
			h.Push(curr)
			i = h.fill(i, experday, ex)
			dayneed++
		} else {
			ik++
		}
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, experday, i int
	var ex []int
	var h Heap
	fmt.Println("First two")
	fmt.Scan(&n, &experday)
	fmt.Println("Enter ex")
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(r, &x)
		ex = append(ex, x)
	}
	fmt.Println(h.findMinIK(experday, i, ex))
}
