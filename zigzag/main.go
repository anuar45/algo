package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}

func ZigZagCon(s string, n int) string {
	var r string
	var i, j, k int

	for i = n + 1; i > 0; i = i - 2 {
		for k = j; k < len(s); {
			r += string(s[k])
			k += i
		}
	}

	return r
}
