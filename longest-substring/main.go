package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("ub"))
	fmt.Println(lengthOfLongestSubstring("abcdefghklljkfwkfjwklajewuifhkljf"))

	fmt.Println(isStrCharsUnique("ub"))
	fmt.Println(isStrCharsUnique("abcdefgh"))
	fmt.Println(isStrCharsUnique("u"))
	fmt.Println(isStrCharsUnique("ub"))
	fmt.Println(isStrCharsUnique("abcdefghh"))
	fmt.Println(isStrCharsUnique("abcddefgh"))
}

func lengthOfLongestSubstring(s string) int {
	var uniqueSubStr string

	if len(s) == 1 {
		return 1
	}
	if len(s) == 0 {
		return 0
	}

	for i := 0; i < len(s); i++ {
		for k := i; k < len(s); k++ {
			if isStrCharsUnique(s[i:k]) && len(s[i:k]) > len(uniqueSubStr) {
				fmt.Println(i, k, s[i:k])
				uniqueSubStr = s[i:k]
				fmt.Println(uniqueSubStr)
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
