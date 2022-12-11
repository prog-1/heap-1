package main

import "fmt"

func main() {

	var n, t int
	fmt.Scanf("%v%v", &n, &t)
	book := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&book[i])
	}

	//t := 4
	//book := []int{28, 59, 87, 19, 19, 17, 69, 72}

	//book := []int{442, 442, 443, 444, 444, 444, 445, 451, 453, 455, 456, 459, 461, 461, 465, 480, 480, 481, 481, 482, 482, 488, 489, 491, 493, 494, 496, 497, 499, 500, 500, 500}
	//t := 93

	fmt.Println(ic(book, t))
}

func ic(book []int, t int) (int, int) {
	var ic, mic, nt, st, days int
	var h Heap
	book2 := book

	if t > len(book) {
		t = len(book)
	}

	for len(book) != 0 || len(h) != 0 {
		for len(h) != 0 && ic >= h[0] {
			h.Pop()
			ic++
			st++
		}
		if len(h) < t && len(book) != 0 {
			nt = t - len(h)
			if nt > len(book) {
				nt = len(book)
			}
			for i := 0; i < nt; i++ { // len(book) != 0 //i < len(book)
				h.Push(book[i])
			}
			book = book[nt:]
		} else if len(h) != 0 {
			ic = h[0]
			mic = ic - st
		}
	}

	ic = mic
	nt = 0
	for len(book2) != 0 || len(h) != 0 {

		if len(h) < t && len(book2) != 0 {
			nt = t - len(h)
			for i := 0; i < nt && i < len(book2); i++ {
				h.Push(book2[i])
			}
			if nt > len(book2) {
				nt = len(book2)
			}
			book2 = book2[nt:]
		}
		for len(h) != 0 && ic >= h[0] {
			h.Pop()
			ic++
			st++
		}
		days++
	}

	return mic, days
}
