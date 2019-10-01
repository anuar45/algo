package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}

func ZigZagCon(s string, n int) string {
	var r string

	for i := n; i >= 0; i-- {
		for k := 0; k <= len(r); {
			r += string(s[k])
			k += i
		}
	}

	return r
}
