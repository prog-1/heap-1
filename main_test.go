package main

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	for _, tc := range []struct {
		name string
		h    Heap
		x    int
		want Heap
	}{
		{"test-1", nil, 1, Heap{1}},
		{"test-2", Heap{1}, 2, Heap{1, 2}},
		{"test-3", Heap{1, 2, 3}, 2, Heap{1, 2, 3, 2}},
		{"test-4", Heap{1, 2, 3, 2}, 2, Heap{1, 2, 3, 2, 2}},
		{"test-5", Heap{1, 2, 3, 4, 5}, 2, Heap{1, 2, 2, 4, 5, 3}},
		{"test-6", Heap{1, 3, 9, 7, 5}, 4, Heap{1, 3, 4, 7, 5, 9}},
		{"test-7", Heap{2, 3, 2, 4}, 1, Heap{1, 2, 2, 4, 3}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tc.h.Push(tc.x)
			if !reflect.DeepEqual(tc.h, tc.want) {
				t.Errorf("got = %v, want = %v", tc.h, tc.want)
			}
		})
	}
}

func TestPop(t *testing.T) {
	for _, tc := range []struct {
		name  string
		h     Heap
		wantE int
		wantH Heap
	}{
		{"case-1", Heap{}, 0, Heap{}},
		{"case-2", Heap{1}, 1, Heap{}},
		{"case-3", Heap{1, 2}, 1, Heap{2}},
		{"case-4", Heap{1, 2, 3}, 1, Heap{2, 3}},
		{"case-5", Heap{1, 2, 3, 4}, 1, Heap{2, 3, 4}},
		{"case-6", Heap{1, 3, 4, 7, 5, 9}, 1, Heap{3, 4, 5, 7, 9}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			x := tc.h.Pop()
			if x != tc.wantE {
				t.Errorf("got = %v, want = %v", x, tc.wantE)
			}
			if !reflect.DeepEqual(tc.h, tc.wantH) {
				t.Errorf("got = %v, want = %v", tc.h, tc.wantH)
			}
		})
	}
}
