package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Run thru tests")
}

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
