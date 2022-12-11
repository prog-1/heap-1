package main

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	for _, tc := range []struct {
		name  string
		h     Heap
		input []int
		want  Heap
	}{
		{"nil", nil, nil, nil},
		{"1", []int{1, 3, 9, 7, 5}, []int{4}, []int{1, 3, 4, 7, 5, 9}},
	} {
		for _, el := range tc.input {
			tc.h.Push(el)
		}
		if !reflect.DeepEqual(tc.h, tc.want) {
			t.Errorf("got = %v, want = %v", tc.h, tc.want)
		}
	}
}

func TestPop(t *testing.T) {
	for _, tc := range []struct {
		name     string
		h        Heap
		wantMin  int
		wantHeap Heap
	}{
		{"nil", nil, 0, nil},
		{"1", []int{1, 3, 4, 7, 5, 9}, 1, []int{3, 5, 4, 7, 9}},
	} {
		if got := tc.h.Pop(); got != tc.wantMin {
			t.Errorf("Incorrect output: got = %v, want = %v", got, tc.wantMin)
		}
		if !reflect.DeepEqual(tc.h, tc.wantHeap) {
			t.Errorf("Incorrect heap order: got = %v, want = %v", tc.h, tc.wantHeap)
		}
	}
}
