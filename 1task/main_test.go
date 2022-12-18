package main

import (
	"testing"
)

func TestPush(t *testing.T) {
	for _, tc := range []struct {
		i    int
		h    Heap
		want []int
	}{
		{0, nil, []int{0}},
		{1, nil, []int{1}},
		{4, []int{1, 3, 9, 7, 5}, []int{1, 3, 4, 7, 5, 9}},
		{6, []int{1, 3, 4, 7, 5, 9}, []int{1, 3, 4, 7, 5, 9, 6}},
		{3, []int{1, 3, 4, 7, 5, 9}, []int{1, 3, 3, 7, 5, 9, 4}},
		{12, []int{1, 3, 3, 7, 5, 9, 4}, []int{1, 3, 3, 7, 5, 9, 4, 12}},
		{6, []int{1, 3, 3, 7, 5, 9, 4, 12}, []int{1, 3, 3, 6, 5, 9, 4, 12, 7}},
	} {
		tc.h.Push(tc.i)
		got := make([]int, len(tc.h))
		copy(got, tc.h)
		if !equal(got, tc.want) {
			t.Errorf("Test(%v) = %v, want = %v", tc.h, got, tc.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestPop(t *testing.T) {
	for _, tc := range []struct {
		h     Heap
		want1 int
		want  []int
	}{
		{[]int{1, 3, 9, 7, 5}, 1, []int{3, 5, 9, 7}},
		{[]int{1, 3, 4, 7, 5, 9}, 1, []int{3, 5, 4, 7, 9}},
		{[]int{1, 3, 3, 7, 5, 9, 4}, 1, []int{3, 3, 4, 7, 5, 9}},
		{[]int{1, 3, 3, 7, 5, 9, 4, 12}, 1, []int{3, 3, 4, 7, 5, 9, 12}},
	} {
		got1 := tc.h.Pop()
		got := make([]int, len(tc.h))
		copy(got, tc.h)
		if !equal(got, tc.want) || got1 != tc.want1 {
			t.Errorf("Test(%v) = %v, %v, want = %v", tc.h, got1, got, tc.want)
		}
	}
}
