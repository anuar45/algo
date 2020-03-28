package main

import (
	"testing"
)

type TestCase struct {
	A []int
	B []int
	R float64
}

var (
	testCases = []TestCase{
		{
			[]int{3, 1, 8, 9},
			[]int{4, 6, 9, 7},
			6.5,
		},
		{
			[]int{1, 2},
			[]int{3, 4},
			2.5,
		},
	}
)

func TestFindMedianSortedArraysDumb(t *testing.T) {
	for i, tc := range testCases {
		got := findMedianSortedArraysDumb(tc.A, tc.B)
		want := tc.R
		if got != want {
			t.Errorf("Error in test case number %d\nInput: %v\n Want: %f\nGot: %f", i, tc, want, got)
		}
	}
}

func TestFindMedianSortedArraysRecursive(t *testing.T) {
	for i, tc := range testCases {
		got := findMedianSortedArraysRecursive(tc.A, tc.B)
		want := tc.R
		if got != want {
			t.Errorf("Error in test case number %d\nInput: %v\n Want: %f\nGot: %f", i, tc, want, got)
		}
	}
}

func TestFindMedianSortedArraysQuick(t *testing.T) {
	for i, tc := range testCases {
		got := findMedianSortedArraysQuick(tc.A, tc.B)
		want := tc.R
		if got != want {
			t.Errorf("Error in test case number %d\nInput: %v\n Want: %f\nGot: %f", i, tc, want, got)
		}
	}
}
