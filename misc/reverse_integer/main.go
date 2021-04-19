package main

import "fmt"

func main() {
	fmt.Println(reverse(325))
	fmt.Println(reverse(-325))
	fmt.Println(reverse(1563847412))
	fmt.Println(reverse(-1563847412))

	fmt.Println((1 << 31) - 1)
	fmt.Println(-1 << 31)
}

func reverse(x int) int {
	var pop, r int

	for x > 0 {
		pop = x % 10
		x /= 10

		r = r*10 + pop
	}

	for x < 0 {
		pop = x % 10
		x /= 10

		r = r*10 + pop
	}

	if r > (1<<31)-1 || r < -1<<31 {
		return 0
	}

	return r
}
