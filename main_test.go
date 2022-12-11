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
		{"1", []int{1, 2, 4, 9, 3, 8, 7, 10, 16, 14}, []int{1}, []int{1, 1, 4, 9, 2, 8, 7, 10, 16, 14, 3}},
		{"2", []int{1, 2, 4, 9, 3, 8, 7, 10, 16, 14}, []int{1, 2, 3}, []int{1, 1, 2, 9, 2, 3, 7, 10, 16, 14, 3, 8, 4}},
		{"3", []int{2, 4, 3}, []int{1, 2, 3}, []int{1, 2, 3, 4, 2, 3}},
		{"4", []int{1, 3, 9, 7, 5}, []int{4}, []int{1, 3, 4, 7, 5, 9}},
		{"5", []int{2}, []int{1}, []int{1, 2}},
		{"6", []int{}, []int{1}, []int{1}},
		{"7", []int{}, []int{}, []int{}},
		{"8", nil, nil, nil},
		// (c) My brotha :)
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
		{"1", []int{1, 3, 4, 7, 5, 9}, 1, []int{3, 5, 4, 7, 9}},
		{"2", []int{1, 2, 4, 9, 3, 8, 7, 10, 16, 14}, 1, []int{2, 3, 4, 9, 14, 8, 7, 10, 16}},
		{"3", []int{1, 2, 3}, 1, []int{2, 3}},
		{"4", []int{1, 3}, 1, []int{3}},
		{"5", []int{2}, 2, []int{}},
		//{"6", []int{}, 0, []int{}},
		//{"7", nil, 0, nil},
		// (c) My brotha as well
	} {
		if got := tc.h.Pop(); got != tc.wantMin {
			t.Errorf("Incorrect output: got = %v, want = %v", got, tc.wantMin)
		}
		if !reflect.DeepEqual(tc.h, tc.wantHeap) {
			t.Errorf("Incorrect heap order: got = %v, want = %v", tc.h, tc.wantHeap)
		}
	}
}

func TestIngusCoefficient(t *testing.T) {
	type Input struct {
		complexities []int
		mit          int
	}
	type Want struct {
		miic, days int
	}
	for _, tc := range []struct {
		name  string
		input Input
		want  Want
	}{
		{"1", Input{[]int{3, 1, 7, 2, 6, 3}, 3}, Want{2, 3}},
		{"2", Input{[]int{7, 6, 5, 4, 3}, 9}, Want{3, 1}},
	} {
		gotMiic, gotDays := IngusCoefficient(tc.input.complexities, tc.input.mit)
		if gotMiic != tc.want.miic {
			t.Errorf("Incorrect minimal initial Ingus coefficient: got = %v, want = %v", gotMiic, tc.want.miic)
		}
		if gotDays != tc.want.days {
			t.Errorf("Incorrect day count: got = %v, want = %v", gotDays, tc.want.days)
		}
	}
}
