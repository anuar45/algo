package main

import (
	"testing"
)

type TestCase struct {
	Data   string
	Result string
}

var (
	testCases = []TestCase{
		TestCase{"fwiofjwebababoripowef", "babab"},
		TestCase{"b", "b"},
		TestCase{"bb", "bb"},
	}
)

func TestLongestPalindrome(t *testing.T) {
	for i, tc := range testCases {
		got := longestPalindrome(tc.Data)
		want := tc.Result
		if got != want {
			t.Errorf("Test %d - Incorrect result.\nWant: %v\nGot: %v\n", i, want, got)
		}
	}
}

func TestReverseStringQuick(t *testing.T) {
	s := "aba"
	got := reverseStringQuick(s)
	want := s
	if got != want {
		t.Errorf("Reverse string failed:\nWant: %s\nGot: %s", want, got)
	}
}
