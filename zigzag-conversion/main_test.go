package main

import (
	"testing"
)

type TestCase struct {
	input string
	rows int
	result string
}

var (
	testCases = []TestCase{
		{
			"abcdefgh",
			3,
			"""
a e g
bdf 
c
""",
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
