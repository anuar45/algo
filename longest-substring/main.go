package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func lengthOfLongestSubstring(s string) int {
	var uniqueSubStr string

	if len(s) == 1 {
		return 1
	}

	for i := 0; i < len(s); i++ {
		for k := i; k < len(s); k++ {
			if isStrCharsUnique(s[i:k]) && len(uniqueSubStr) < len(s[i:k]) {
				uniqueSubStr = s[i:k]
			}
		}
	}

	return len(uniqueSubStr)
}

func isStrCharsUnique(s string) bool {
	for i := 0; i < len(s); i++ {
		for k := i + 1; k < len(s); k++ {
			if s[i] == s[k] {
				return false
			}
		}
	}
	return true
}
