package main

type Heap []int

func Parent(i int) int {
	return (i - 1) / 2
}

func Left(i int) int {
	return 2*i + 1
}
func Right(i int) int {
	return 2*i + 2
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Push(x int) {
	*h = append(*h, x)
	for i := len(*h) - 1; i > 0 && (*h)[i] < (*h)[Parent(i)]; i = Parent(i) {
		(*h)[i], (*h)[Parent(i)] = (*h)[Parent(i)], (*h)[i]
	}
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Pop() int {
	i := len(*h) - 1
	if len(*h) == 0 {
		return 0
	}

	(*h)[0], (*h)[i] = (*h)[i], (*h)[0]
	*h = (*h)[:i]

	for i := 0; ; {
		left, right := Left(i), Right(i)
		min := i
		if left < len(*h) && (*h)[i] > (*h)[left] {
			min = left
		}
		if right < len(*h) && (*h)[left] > (*h)[right] {
			min = right
		}
		if min != i {
			(*h)[i], (*h)[min] = (*h)[min], (*h)[i]
			i = min
			continue
		}
		break
	}
	return (*h)[0]
}

func main() {

}
