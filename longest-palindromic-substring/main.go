package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Run thru tests")
}

// This one is bruteforce, doesnt pass leetcode
func longestPalindrome(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(s[i:j+1]) && (len(s[i:j+1]) > len(result)) {
				result = s[i : j+1]
			}
		}
	}

	return result
}

func isPalindrome(s string) bool {
	var result bool
	if s == reverseStringQuick(s) {
		result = true
	}
	return result
}

func reverseStringQuick(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

// This method use expand algorithm
func longestPalindromeQuick(s string) string {
	var left, right int
	var low, high int

	if len(s) == 1 || s == "" || (len(s) == 2 && s[0] == s[1]) {
		return s
	}

	for i := 1; i < len(s); i++ {

		left = i - 1
		right = i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}

		if right-left > high-low {
			low = left
			high = right
		}

		left = i - 1
		right = i + 1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}

		if right-left > high-low {
			low = left
			high = right
		}

	}

	return s[low+1 : high]

}
