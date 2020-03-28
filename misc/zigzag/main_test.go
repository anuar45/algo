package main

import (
	"testing"
)

type TestCase struct {
	A string
	B int
	R string
}

var (
	testCases = []TestCase{
		{
			"PAYPALISHIRING",
			3,
			"PAHNAPLSIIGYIR",
		},
	}
)

func TestZigZagCon(t *testing.T) {
	for i, tc := range testCases {
		got := ZigZagCon(tc.A, tc.B)
		want := tc.R
		if got != want {
			t.Errorf("Error in test case number %d\nTestCaseData: %v\n Want: %s\nGot: %s", i, tc, want, got)
		}
	}
}
