package main

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	for _, tc := range []struct {
		name string
		h    Heap
		n    []int // n = numbers to push
		want []int
	}{
		{"1", []int{1, 2, 4, 9, 3, 8, 7, 10, 16, 14}, []int{1}, []int{1, 1, 4, 9, 2, 8, 7, 10, 16, 14, 3}},
		{"2", []int{1, 2, 4, 9, 3, 8, 7, 10, 16, 14}, []int{1, 2, 3}, []int{1, 1, 2, 9, 2, 3, 7, 10, 16, 14, 3, 8, 4}},
		{"3", []int{2, 4, 3}, []int{1, 2, 3}, []int{1, 2, 3, 4, 2, 3}},
		{"4", []int{1, 3, 9, 7, 5}, []int{4}, []int{1, 3, 4, 7, 5, 9}},
		{"5", []int{2}, []int{1}, []int{1, 2}},
		{"6", []int{}, []int{1}, []int{1}},
		{"7", []int{}, []int{}, []int{}},
		{"8", nil, nil, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			h := tc.h
			for _, el := range tc.n {
				h.Push(el)
			}
			if !reflect.DeepEqual([]int(h), tc.want) {
				t.Errorf("got = %v, want = %v", h, tc.want)
			}
		})
	}
}

func TestPop(t *testing.T) {
	for _, tc := range []struct {
		name string
		h    Heap
		want []int
	}{
		{"1", []int{1, 2, 4, 9, 3, 8, 7, 10, 16, 14}, []int{2, 3, 4, 9, 14, 8, 7, 10, 16}},
		{"2", []int{1, 3, 4, 7, 5, 9}, []int{3, 5, 4, 7, 9}},
		{"3", []int{1, 2, 3}, []int{2, 3}},
		{"4", []int{1, 3}, []int{3}},
		{"5", []int{2}, []int{}},
		{"6", []int{}, []int{}},
		{"7", nil, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			h := tc.h
			h.Pop()
			if !reflect.DeepEqual([]int(h), tc.want) {
				t.Errorf("got = %v, want = %v", h, tc.want)
			}
		})
	}
}
