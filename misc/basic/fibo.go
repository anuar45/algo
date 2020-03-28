package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	var start time.Time

	start = time.Now()
	fmt.Println(fibo(n), time.Since(start), "Recur")

	start = time.Now()
	fmt.Println(fiboFor(n), time.Since(start), "For")

}

func fiboFor(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func fibo(n int) int {
	if n == 1 {
		return 1
	} else if n == 0 {
		return 0
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}
