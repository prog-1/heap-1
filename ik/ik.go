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
	// INPUT
	var taskNum, tasksPerDay int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &taskNum, &tasksPerDay)
	tasks := make([]int, taskNum)
	for i := 0; i < taskNum; i++ {
		var tmp int
		fmt.Fscan(reader, &tmp)
		tasks[i] = tmp
	}
	if tasksPerDay > taskNum {
		tasksPerDay = taskNum
	}
	IK := findIngusCofficent(tasks, tasksPerDay)
	fmt.Println(IK, calculateDays(IK, tasks, tasksPerDay))
}

func findIngusCofficent(tasks []int, tasksPerDay int) int {
	var ik, startIK, i int
	var BST Heap
	for ; i < tasksPerDay; i++ {
		BST.Push(tasks[i])
	}
	for len(BST) != 0 {

		if d := BST.Pop() - ik; d > 0 {
			startIK += d
			ik += d
		}
		if i < len(tasks) {
			BST.Push(tasks[i])
			i++
		}
		ik++
	}
	return startIK
}

func calculateDays(cofficent int, tasks []int, tasksPerDay int) int {

	//Init
	var day, i int
	var BST Heap
	var removed int
	for ; i < tasksPerDay; i++ {
		BST.Push(tasks[i])
	}
	// Solve
	for i < len(tasks) {
		if len(BST) == 0 || BST.Pop() > cofficent {
			day++
			addNewTasks(&i, tasks, &removed, &BST)
			continue
		}
		cofficent++
		removed++
	}

	return day + 1
}

func addNewTasks(i *int, tasks []int, removed *int, BST *Heap) {
	for ; *removed != 0 && *i < len(tasks); *removed-- {
		BST.Push(tasks[*i])
		*i++
	}
}
