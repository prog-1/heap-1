package main

type Heap []int

// func main() {
// 	h := heap{heap: []int{1, 2, 4, 9, 3, 8, 7, 10, 16, 14}}
// 	//h := heap{heap: []int{3, 2, 1}}
// 	h.push(6)
// 	h.push(3)
// 	h.pop()
// 	fmt.Println(h)

// }

//heap left child
func l(i int) int {
	return 2*i + 1
}

//heap right child
func r(i int) int {
	return 2*i + 2
}

//heap parent
func p(i int) int {
	return (i - 1) / 2
}

func (h *Heap) Push(val int) {
	(*h) = append((*h), val)                   //add this element at the end of the heap
	el := len((*h)) - 1                        //storing cur el index
	for ; (*h)[el] < (*h)[p(el)]; el = p(el) { //checking is cur val < parent val
		(*h)[el], (*h)[p(el)] = (*h)[p(el)], (*h)[el] //swapping
	}
}

func (h *Heap) Pop() int {
	if len((*h)) == 0 { // edge case
		return 0
	}
	min := (*h)[0]              //saving min el
	(*h)[0] = (*h)[len((*h))-1] //changing min el on last el
	(*h) = (*h)[:len((*h))-1]   //deleting last el
	h.heapifyup()               //balance the tree
	return min
}

//balance the tree further from [0]
func (h *Heap) heapifyup() {
	for el := 0; h.isvalid(l(el)) && (*h)[el] > (*h)[l(el)] || h.isvalid(r(el)) && (*h)[el] > (*h)[r(el)]; { // till we haven't reached the bottom of the tree or parent is bigger than one of it's childs
		if h.isvalid(l(el)) && !h.isvalid(r(el)) && (*h)[l(el)] < (*h)[el] { //if we have only left child and left child < parent
			(*h)[el], (*h)[l(el)] = (*h)[l(el)], (*h)[el] // swap parent and left child
			el = l(el)
		} else if !h.isvalid(l(el)) && h.isvalid(r(el)) && (*h)[r(el)] < (*h)[el] { //if we have only right child and right child < parent
			(*h)[el], (*h)[r(el)] = (*h)[r(el)], (*h)[el] // swap parent and right child
			el = r(el)
		} else if (*h)[l(el)] < (*h)[r(el)] { //if left child < right child
			(*h)[el], (*h)[l(el)] = (*h)[l(el)], (*h)[el] // swap parent and left child
			el = l(el)
		} else { // right child < left child
			(*h)[el], (*h)[r(el)] = (*h)[r(el)], (*h)[el] // swap parent and right child
			el = r(el)
		}
	}
}

func (h *Heap) isvalid(el int) bool {
	return el >= 0 && el < len((*h)) //if el goes outside of heap borders
}
