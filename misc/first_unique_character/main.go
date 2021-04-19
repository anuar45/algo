package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("leetcode"))
}

func firstUniqChar(s string) int {
	m := make(map[rune]int)

	for _, v := range []rune(s) {
		m[v]++
	}

	for i, v := range []rune(s) {
		if m[v] == 1 {
			return i
		}
	}

	return -1
}
